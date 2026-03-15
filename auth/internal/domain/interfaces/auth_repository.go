package interfaces

import (
	"auth/internal/domain/entity"
)

type AuthRepository interface {
	GenerateToken(tokenType entity.TokenType, user string) (*entity.JWToken, error)
	SaveToken(tkn *entity.JWToken) error
	GetToken(id string) (*entity.JWToken, error)
	CheckTokenExisting(tokenID string) (bool, error)
	ValidateToken(tokenString string) (*entity.JWToken, error)
	RevokeToken(tokenID string) error
}
