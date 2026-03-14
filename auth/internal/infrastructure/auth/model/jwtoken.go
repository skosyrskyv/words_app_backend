package model

import (
	"auth/internal/domain/entity"
	"fmt"
	"time"
)

type JWToken struct {
	TokenID   string    `gorm:"primaryKey"`
	Token     string    `gorm:"type:text;not null"`
	UserUUID  string    `gorm:"index;not null"`
	IssuedAt  time.Time `gorm:"not null"`
	ExpiresAt time.Time `gorm:"index;not null"`
	Type      string    `gorm:"type:varchar(10);not null"`
}

func (JWToken) TableName() string {
	return "jwt_tokens"
}

func (t *JWToken) ToEntity() (*entity.JWToken, error) {
	tokenType, err := entity.ParseTokenType(t.Type)

	if err != nil {
		fmt.Printf("failed to parse token type, error: %s", err.Error())
		tokenType = ""
	}

	return &entity.JWToken{
		TokenID:   t.TokenID,
		Token:     t.Token,
		UserUUID:  t.UserUUID,
		IssuedAt:  t.IssuedAt,
		ExpiresAt: t.ExpiresAt,
		Type:      tokenType,
	}, nil
}

func (t *JWToken) FromEntity(entity entity.JWToken) {
	t.TokenID = entity.TokenID
	t.Token = entity.Token
	t.UserUUID = entity.UserUUID
	t.IssuedAt = entity.IssuedAt
	t.ExpiresAt = entity.ExpiresAt
	t.Type = entity.Type.String()
}
