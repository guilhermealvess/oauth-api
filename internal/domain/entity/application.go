package entity

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Application struct {
	ID          uuid.UUID `validate:"required"`
	TeamID      uuid.UUID `validate:"required"`
	Name        string    `validate:"required"`
	Slug        string    `validate:"required"`
	Description string
	IsActive    bool      `validate:"required"`
	CreatedBy   string    `validate:"required"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required"`
	Permissions []*Permission
}

func NewApplication(name, desc string, user User) (*Application, error) {
	if user.Ok() != nil {
		return nil, errors.New("User is invalid")
	}

	now := time.Now().UTC()
	name = removeSpecialCharacters(name)
	slug := slugFy(name)
	app := Application{
		ID:          uuid.New(),
		TeamID:      user.TeamID,
		Name:        name,
		Description: desc,
		Slug:        slug,
		IsActive:    true,
		CreatedBy:   user.Email,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return &app, app.Ok()
}

func (a *Application) Ok() error {
	return validator.New().Struct(a)
}
