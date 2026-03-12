package interfaces

import "auth/internal/domain/entity"

type AuthRepository interface {
	GenerateAccessToken(user *entity.User) (*entity.JWToken, error)
	GenerateRefreshToken() (*entity.JWToken, error)
	SaveRefreshToken(tkn *entity.JWToken) error
	GetToken(id string) (string, error)
}
