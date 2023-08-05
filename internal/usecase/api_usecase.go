package usecase

import (
	"context"

	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
	"github.com/guilhermealvess/oauth-api/internal/domain/repository"
)

type ApiUsecase interface {
	CreateApiServer(ctx context.Context, input CreateApiServerInput) (*CreateApiServerOutput, error)
}

type apiUse struct {
	apiRepository repository.ApiServerRepository
}

func NewApiUse(repo repository.ApiServerRepository) ApiUsecase {
	return apiUse{repo}
}

func (a apiUse) CreateApiServer(ctx context.Context, input CreateApiServerInput) (*CreateApiServerOutput, error) {
	apiServer, err := entity.NewApiServer(input.Name, input.Description, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	var permissions []*CreatePermissionOutput
	for _, p := range input.Permissions {
		permission, err := entity.NewPermission(apiServer, p.Resource, p.Action)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, &CreatePermissionOutput{
			ID:          permission.ID,
			ApiID:       permission.ApiID,
			ApiResource: permission.ApiResource,
			Action:      permission.Action,
			IsActive:    permission.IsActive,
			CreatedAt:   permission.CreatedAt,
			UpdatedAt:   permission.UpdatedAt,
		})
	}

	if err := a.apiRepository.Save(ctx, apiServer); err != nil {
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
		Permissions: permissions,
	}, nil
}
