package repository

import (
	"context"

	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type ApiServerRepository interface {
	Save(context.Context, *entity.ApiServer) error
}
