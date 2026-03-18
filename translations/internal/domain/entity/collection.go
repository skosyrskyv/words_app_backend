package entity

import "github.com/google/uuid"

type Collection struct {
	UUID uuid.UUID
	User uuid.UUID
	Name string
}
