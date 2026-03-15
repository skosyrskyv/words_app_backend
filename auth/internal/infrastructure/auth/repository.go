package auth

import (
	"auth/config"
	"auth/internal/domain/entity"
	"auth/internal/infrastructure/auth/model"
	"auth/pkg/postgres"
	"auth/pkg/redis"
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type repository struct {
	cfg        config.JWTConfig
	redis      *redis.Redis
	postgres   *postgres.Postgres
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRepository(cfg config.JWTConfig, redis *redis.Redis, postgres *postgres.Postgres) *repository {
	privateKeysBytes, err := os.ReadFile(cfg.PrivateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key, error: %s", err.Error())
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeysBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key, error: %s", err.Error())
	}

	publicKeyBytes, err := os.ReadFile(cfg.PublicKeyPath)
	if err != nil {
		log.Fatalf("Failed to read public key, error: %s", err.Error())
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse public key, error: %s", err.Error())
	}

	return &repository{
		cfg:        cfg,
		redis:      redis,
		postgres:   postgres,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (r *repository) GenerateToken(tokenType entity.TokenType, user string) (*entity.JWToken, error) {
	tokenId, err := generateTokenId()
	if err != nil {
		return nil, err
	}

	var ttl time.Duration

	switch tokenType {
	case entity.AccessTokenType:
		ttl, err = r.cfg.GetAccessTTL()
	case entity.RefreshTokenType:
		ttl, err = r.cfg.GetRefreshTTL()
	default:
		return nil, fmt.Errorf("invalid token type: %s", tokenType)
	}

	if err != nil {
		return nil, err
	}

	iat := time.Now()
	exp := time.Now().Add(ttl)

	claims := jwt.MapClaims{
		"jti": tokenId,
		"sub": user,
		"iss": r.cfg.Issuer,
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenSigned, err := token.SignedString(r.privateKey)
	if err != nil {
		return nil, err
	}

	model := &model.JWToken{
		TokenID:   tokenId,
		IssuedAt:  iat,
		UserUUID:  user,
		ExpiresAt: exp,
		Token:     tokenSigned,
		Type:      string(tokenType),
	}

	return model.ToEntity()
}

func (r *repository) SaveToken(tkn *entity.JWToken) error {
	var model model.JWToken
	model.FromEntity(*tkn)

	if model.Type == entity.RefreshTokenType.String() {
		err := r.redis.Set(context.Background(), model.TokenID, model.Token, time.Hour*24*7)
		if err != nil {
			slog.Error("Failed to save token in redis", "error", err.Error())
		}
	}

	if err := r.postgres.DB().Save(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetToken(id string) (*entity.JWToken, error) {
	token := r.postgres.DB().Where("token_id = ?", id).First(&model.JWToken{})
	if token.Error != nil {
		return nil, token.Error
	}
	var entityToken entity.JWToken
	if res := token.Scan(&entityToken); res.Error != nil {
		return nil, res.Error
	}
	return &entityToken, nil
}

func (r *repository) RevokeToken(tokenID string) error {
	if err := r.redis.Del(context.Background(), tokenID); err != nil {
		return err
	}

	if err := r.postgres.DB().Where("token_id = ?", tokenID).Delete(&model.JWToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) CheckTokenExisting(tokenID string) (bool, error) {
	token, err := r.redis.Get(context.Background(), tokenID)
	if err != nil || token == "" {
		res := r.postgres.DB().Where("token_id = ?", tokenID).First(&model.JWToken{})
		if res.Error != nil {
			return false, res.Error
		}
		return false, fmt.Errorf("Token not found, error: %s", err.Error())
	}

	return true, nil
}

func (r *repository) ValidateToken(tokenString string) (*entity.JWToken, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return r.publicKey, nil
		},
		jwt.WithIssuer(r.cfg.Issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	tokenID, ok := claims["jti"].(string)
	if !ok || tokenID == "" {
		return nil, fmt.Errorf("missing jti claim")
	}

	userUUID, err := claims.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("missing sub claim: %w", err)
	}

	iat, err := claims.GetIssuedAt()
	if err != nil {
		return nil, fmt.Errorf("missing iat claim: %w", err)
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("missing exp claim: %w", err)
	}

	return &entity.JWToken{
		TokenID:   tokenID,
		Token:     tokenString,
		UserUUID:  userUUID,
		IssuedAt:  iat.Time,
		ExpiresAt: exp.Time,
	}, nil
}

func generateTokenId() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
