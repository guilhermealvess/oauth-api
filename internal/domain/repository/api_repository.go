package repository

import "github.com/guilhermealvess/oauth-api/internal/domain/entity"

type ApiServerRepository interface {
	Save(*entity.ApiServer) error
}
