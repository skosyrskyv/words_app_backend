package interfaces

import "auth/internal/domain/entity"

type AuthRepository interface {
	GenerateToken(tokenType entity.TokenType, user *entity.User) (*entity.JWToken, error)
	SaveToken(tkn *entity.JWToken) error
	GetToken(id string) (*entity.JWToken, error)
}
