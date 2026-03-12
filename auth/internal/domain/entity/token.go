package entity

import "time"

type JWToken struct {
	TokenID   string
	Token     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
