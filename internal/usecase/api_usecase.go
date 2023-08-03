package usecase

import (
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
	"github.com/guilhermealvess/oauth-api/internal/domain/repository"
)

type ApiUsecase interface {
	CreateApiServer(input CreateApiServerInput) (*CreateApiServerOutput, error)
}

type apiUse struct {
	apiRepository repository.ApiServerRepository
}

func NewApiUse(repo repository.ApiServerRepository) ApiUsecase {
	return apiUse{repo}
}

func (a apiUse) CreateApiServer(input CreateApiServerInput) (*CreateApiServerOutput, error) {
	apiServer, err := entity.NewApiServer(input.Name, input.Description, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	for _, p := range input.Permissions {
		permission, err := entity.NewPermission(apiServer, p.Resource, p.Action)
		if err != nil {
			return nil, err
		}
		apiServer.Permissions = append(apiServer.Permissions, permission)
	}

	if err := a.apiRepository.Save(apiServer); err != nil {
		return nil, err
	}

	return &CreateApiServerOutput{
		ID:          apiServer.ID,
		Name:        apiServer.Name,
		Description: apiServer.Description,
		Slug:        apiServer.Slug,
		IsActive:    apiServer.IsActive,
		CreatedBy:   apiServer.CreatedBy,
		CreatedAt:   apiServer.CreatedAt,
		UpdatedAt:   apiServer.UpdatedAt,
		Permissions: apiServer.Permissions,
	}, nil
}
