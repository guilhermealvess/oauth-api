package usecase

import (
	"time"

	"github.com/google/uuid"
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
	Permissions []*CreatePermissionOutput
}

type CreatePermissionOutput struct {
	ID          uuid.UUID `json:"id"`
	ApiID       uuid.UUID `json:"api_id"`
	ApiResource string    `json:"api_resource"`
	Action      string    `json:"action"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
