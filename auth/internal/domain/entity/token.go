package entity

import (
	"fmt"
	"strings"
	"time"
)

type TokenType string

const (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

type JWToken struct {
	TokenID   string
	Token     string
	UserUUID  string
	IssuedAt  time.Time
	ExpiresAt time.Time
	Type      TokenType
}

func (t TokenType) String() string {
	return string(t)
}

func ParseTokenType(str string) (TokenType, error) {
	switch strings.TrimSpace(strings.ToLower(str)) {
	case "access":
		return AccessTokenType, nil
	case "refresh":
		return RefreshTokenType, nil
	default:
		return "", fmt.Errorf("invalid token type: %q", str)
	}
}
