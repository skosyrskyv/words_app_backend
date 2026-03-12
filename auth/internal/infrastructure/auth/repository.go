package auth

import (
	"auth/config"
	"auth/internal/domain/entity"
	"auth/pkg/redis"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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

	ttl, err := r.cfg.GetAccessTLL()

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

	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &entity.JWToken{
		TokenID:   tokenId,
		IssuedAt:  iat,
		ExpiresAt: exp,
		Token:     tokenString,
	}, nil
}

func (r *repository) GenerateRefreshToken() (*entity.JWToken, error) {
	tokenId, err := generateTokenId()
	if err != nil {
		return nil, err
	}

	ttl, err := r.cfg.GetRefreshTLL()

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
		log.Fatal("Ошибка чтения приватного ключа:", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatal("Ошибка парсинга приватного ключа:", err)
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &entity.JWToken{
		TokenID:   tokenId,
		IssuedAt:  iat,
		ExpiresAt: exp,
		Token:     tokenString,
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

func parsePrivateKey(keyBytes []byte) (*rsa.PrivateKey, error) {
	// Декодируем PEM
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("не удалось декодировать PEM блок")
	}

	fmt.Printf("Тип PEM блока: %s\n", block.Type)

	// Пробуем разные форматы парсинга
	var key interface{}
	var err error

	// Пробуем PKCS1
	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		fmt.Println("Ключ в формате PKCS1")
		return key.(*rsa.PrivateKey), nil
	}

	// Пробуем PKCS8
	key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		fmt.Println("Ключ в формате PKCS8")
		if rsaKey, ok := key.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, fmt.Errorf("ключ не RSA")
	}

	return nil, fmt.Errorf("неподдерживаемый формат ключа")
}
