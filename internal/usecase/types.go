package usecase

import (
	"time"

	"github.com/google/uuid"
)

/* type ApplicationOutput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Permissions []*CreatePermissionOutput
} */

type CreatePermissionOutput struct {
	ID          uuid.UUID `json:"id"`
	ApiID       uuid.UUID `json:"api_id"`
	ApiResource string    `json:"api_resource"`
	Action      string    `json:"action"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TeamOutput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CreateUserInput struct {
	TeamID                            uuid.UUID
	Name, Occupation, Email, Password string
}

type UserOutput struct {
	ID         uuid.UUID `json:"id"`
	TeamID     uuid.UUID `json:"team_id"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
}

type CreateApplication struct {
	UserEmail   string             `json:"user_email"` // payload JWT
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Permissions []CreatePermission `json:"permissions"`
}

type ApplicationOutput struct {
	ID          uuid.UUID          `json:"id"`
	TeamID      uuid.UUID          `json:"team_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Slug        string             `json:"slug"`
	CreatedBy   string             `json:"created_at"`
	Permissions []PermissionOutput `json:"permissions"`
}

type CreatePermission struct {
	Action    string `json:"action"`
	Resource  string `json:"resource"`
	CreatedBy string `json:"created_by"`
}

type PermissionOutput struct {
	ID            uuid.UUID `json:"id"`
	ApplicationID uuid.UUID `json:"application_id"`
	Action        string    `json:"action"`
	Resource      string    `json:"resource"`
	CreatedBy     string    `json:"createdBy"`
}
