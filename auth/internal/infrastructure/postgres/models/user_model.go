package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Uuid      string         `gorm:"primarykey" json:"uuid"`
	Email     string         `gorm:"uniqueIndex" json:"email" binding:"required,email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
