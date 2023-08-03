package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type CreateApiServerInput struct {
	Name        string `json:"name"`
	Description string `json:"descriptiom"`
	CreatedBy   string `json:"created_by"`
	Permissions []struct {
		Action   string `json:"action"`
		Resource string `json:"resource"`
	} `json:"permissions"`
}

type CreateApiServerOutput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Permissions []*entity.Permission
}
