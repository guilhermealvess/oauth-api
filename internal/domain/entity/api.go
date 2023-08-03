package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ApiServer struct {
	ID          uuid.UUID `validate:"required"`
	Name        string    `validate:"required"`
	Slug        string    `validate:"required"`
	Description string
	IsActive    bool      `validate:"required"`
	CreatedBy   string    `validate:"required"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required"`
	Permissions []*Permission
}

func NewApiServer(name, desc, author string) (*ApiServer, error) {
	now := time.Now().UTC()
	api := ApiServer{
		ID:          uuid.New(),
		Name:        name,
		Description: desc,
		IsActive:    true,
		Slug:        name, // TODO:
		CreatedBy:   author,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return &api, api.Ok()
}

func (a *ApiServer) Ok() error {
	return validator.New().Struct(a)
}
