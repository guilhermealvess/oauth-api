package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Team struct {
	ID          uuid.UUID `validate:"required"`
	Name        string    `validate:"required"`
	Slug        string    `validate:"required"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required"`
	Description string
}

func NewTeam(name, description string) (*Team, error) {
	now := time.Now().UTC()
	name = removeSpecialCharacters(name)
	t := Team{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slugFy(name),
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return &t, t.Ok()
}

func (t *Team) Ok() error {
	return validator.New().Struct(t)
}
