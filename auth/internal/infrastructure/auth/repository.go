package auth

import (
	"auth/config"
	"auth/internal/domain/entity"
	"auth/pkg/redis"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//  SEC-001 — Medium

//   Локация: auth/internal/infrastructure/auth/repository.go:72,105
//   Категория: E — Чтение ключа с диска на каждый запрос

//   Описание: Приватный RSA-ключ читается из файла при каждой генерации токена. Это и I/O-overhead, и потенциальная
//   уязвимость (race condition при замене файла).

//   Рекомендация: Загружать ключ один раз при инициализации repository, хранить в поле структуры.

type repository struct {
	cfg config.JWTConfig
	db  *redis.Redis
}

func NewRepository(cfg config.JWTConfig, db *redis.Redis) *repository {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *repository) GenerateAccessToken(user *entity.User) (*entity.JWToken, error) {
	tokenId, err := generateTokenId()
	if err != nil {
		return nil, err
	}

	ttl, err := r.cfg.GetAccessTTL()

	if err != nil {
		// Log error
	}

	iat := time.Now()
	exp := time.Now().Add(ttl)

	claims := jwt.MapClaims{
		"jti": tokenId,
		"sub": user.UUID,
		"iss": r.cfg.Issuer,
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privateKeyBytes, err := os.ReadFile(r.cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	tokenSigned, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &entity.JWToken{
		TokenID:   tokenId,
		IssuedAt:  iat,
		ExpiresAt: exp,
		Token:     tokenSigned,
	}, nil
}

func (r *repository) GenerateRefreshToken() (*entity.JWToken, error) {
	tokenId, err := generateTokenId()
	if err != nil {
		return nil, err
	}

	ttl, err := r.cfg.GetRefreshTTL()

	if err != nil {
		// Log error
	}

	iat := time.Now()
	exp := time.Now().Add(ttl)

	claims := jwt.MapClaims{
		"jti": tokenId,
		"iss": r.cfg.Issuer,
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKeyBytes, err := os.ReadFile(r.cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	tokenSigned, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &entity.JWToken{
		TokenID:   tokenId,
		IssuedAt:  iat,
		ExpiresAt: exp,
		Token:     tokenSigned,
	}, nil
}

func (r *repository) SaveRefreshToken(tkn *entity.JWToken) error {
	if err := r.db.Set(context.Background(), tkn.TokenID, tkn.Token, 0); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}
	return nil
}

func (r *repository) GetToken(id string) (string, error) {
	return "", nil
}

func generateTokenId() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
