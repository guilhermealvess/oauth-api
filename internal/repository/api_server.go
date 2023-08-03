package repository

import (
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
	"github.com/guilhermealvess/oauth-api/internal/domain/repository"
)

type apiServerRepository struct{}

func NewApiServerRepository() repository.ApiServerRepository {
	return apiServerRepository{}
}

func (a apiServerRepository) Save(apiServer *entity.ApiServer) error {
	return nil
}
