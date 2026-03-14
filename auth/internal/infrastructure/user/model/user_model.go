package model

import (
	"auth/internal/domain/entity"
	"auth/internal/domain/entity/vo"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	Uuid      string         `gorm:"primaryKey;type:uuid;column:uuid"`
	Email     string         `gorm:"uniqueIndex;type:varchar(255);not null;column:email"`
	Password  string         `gorm:"type:varchar(255);not null;column:password"`
	IsActive  bool           `gorm:"default:true;not null;column:is_active"`
	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (UserModel) TableName() string {
	return "users"
}

func (model *UserModel) ToEntity() (*entity.User, error) {
	uid, err := uuid.Parse(model.Uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse UUID: %w", err)
	}

	email, err := vo.NewEmailFromString(model.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create email value object: %w", err)
	}

	passwordHash, err := vo.NewPasswordHashFromString(model.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create password hash value object: %w", err)
	}

	return &entity.User{
		UUID:         uid,
		Email:        email,
		PasswordHash: passwordHash,
		IsActive:     model.IsActive,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
	}, nil
}

func (model *UserModel) FromEntity(entity entity.User) *UserModel {
	return &UserModel{
		Uuid:      entity.UUID.String(),
		Email:     entity.Email.String(),
		Password:  entity.PasswordHash.String(),
		IsActive:  entity.IsActive,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
