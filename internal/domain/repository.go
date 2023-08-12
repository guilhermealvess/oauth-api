package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type TeamRepository interface {
	Save(context.Context, *entity.Team) error
	FindByID(context.Context, uuid.UUID) (*entity.Team, error)
}

type UserRepository interface {
	Save(context.Context, *entity.User) error
	FindByID(context.Context, uuid.UUID) (*entity.User, error)
	FindByEmail(context.Context, string) (*entity.User, error)
}

type ApplicationRepository interface {
	Save(context.Context, *entity.Application) error
}
