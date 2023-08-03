package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID `validate:"required"`
	ApiID       uuid.UUID `validate:"required"`
	ApiResource string    `validate:"required"`
	Action      string    `validate:"required"`
	IsActive    bool      `validate:"required"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required"`
}

func NewPermission(api *ApiServer, resource, action string) (*Permission, error) {
	now := time.Now().UTC()
	permission := Permission{
		ID:          uuid.New(),
		ApiID:       api.ID,
		ApiResource: resource,
		Action:      action,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return &permission, permission.Ok()
}

func (p *Permission) Ok() error {
	return validator.New().Struct(p)
}
