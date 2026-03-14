package auth

import (
	"auth/config"
	"auth/internal/domain/entity"
	"auth/internal/infrastructure/auth/model"
	"auth/pkg/postgres"
	"auth/pkg/redis"
	"crypto/rsa"
	"fmt"
	"log"
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

	return &repository{
		cfg:        cfg,
		redis:      redis,
		postgres:   postgres,
		privateKey: privateKey,
	}
}

func (r *repository) GenerateToken(tokenType entity.TokenType, user *entity.User) (*entity.JWToken, error) {
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
		"sub": user.UUID.String(),
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
		UserUUID:  user.UUID.String(),
		ExpiresAt: exp,
		Token:     tokenSigned,
		Type:      string(tokenType),
	}

	return model.ToEntity()
}

func (r *repository) SaveToken(tkn *entity.JWToken) error {
	var model model.JWToken
	model.FromEntity(*tkn)

	r.postgres.DB().Save(&model)

	return nil
}

func (r *repository) GetToken(id string) (*entity.JWToken, error) {
	return nil, nil
}

func generateTokenId() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
