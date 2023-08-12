package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Permission struct {
	ID                  uuid.UUID `validate:"required"`
	ApplicationID       uuid.UUID `validate:"required"`
	ApplicationResource string    `validate:"required"`
	Action              string    `validate:"required"`
	CreatedBy           string    `validate:"required"`
	CreatedAt           time.Time `validate:"required"`
	UpdatedAt           time.Time `validate:"required"`
}

func NewPermission(application Application, resource, action, createdBy string) (*Permission, error) {
	now := time.Now().UTC()
	permission := Permission{
		ID:                  uuid.New(),
		ApplicationID:       application.ID,
		ApplicationResource: resource,
		Action:              action,
		CreatedBy:           createdBy,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	return &permission, permission.Ok()
}

func (p *Permission) Ok() error {
	return validator.New().Struct(p)
}
