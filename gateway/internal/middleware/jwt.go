package middleware

import (
	"crypto/rsa"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"gateway/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	ContextUserUUID = "user_uuid"
	ContextTokenID  = "token_id"
)

type JWTMiddleware struct {
	publicKey *rsa.PublicKey
	issuer    string
	logger    *slog.Logger
}

func NewJWTMiddleware(cfg config.JWTConfig, logger *slog.Logger) *JWTMiddleware {
	publicKeyBytes, err := os.ReadFile(cfg.PublicKeyPath)
	if err != nil {
		log.Fatalf("JWT middleware: failed to read public key: %s", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatalf("JWT middleware: failed to parse public key: %s", err)
	}

	return &JWTMiddleware{
		publicKey: publicKey,
		issuer:    cfg.Issuer,
		logger:    logger,
	}
}

func (m *JWTMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return m.publicKey, nil
		}, jwt.WithIssuer(m.issuer), jwt.WithExpirationRequired())

		if err != nil || !token.Valid {
			m.logger.Debug("JWT validation failed", slog.Any("error", err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		userUUID, _ := claims["sub"].(string)
		tokenID, _ := claims["jti"].(string)

		// Передаём данные дальше через заголовки к микросервисам
		c.Request.Header.Set("X-User-UUID", userUUID)
		c.Request.Header.Set("X-Token-ID", tokenID)

		c.Set(ContextUserUUID, userUUID)
		c.Set(ContextTokenID, tokenID)

		c.Next()
	}
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", jwt.ErrTokenUnverifiable
	}
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return "", jwt.ErrTokenMalformed
	}
	return parts[1], nil
}
