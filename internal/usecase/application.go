package usecase

import (
	"context"

	"github.com/guilhermealvess/oauth-api/internal/domain"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type ApplicationUseCase interface {
	CreateApplication(ctx context.Context, input CreateApplication) (*ApplicationOutput, error)
}

type applicationUse struct {
	repository     domain.ApplicationRepository
	userRepository domain.UserRepository
}

func NewApplicationUseCase(repo domain.ApplicationRepository, userRepository domain.UserRepository) ApplicationUseCase {
	return applicationUse{repo, userRepository}
}

func (a applicationUse) CreateApplication(ctx context.Context, input CreateApplication) (*ApplicationOutput, error) {
	user, err := a.userRepository.FindByEmail(ctx, input.UserEmail)
	if err != nil {
		return nil, err
	}

	if err := user.Ok(); err != nil {
		return nil, err
	}

	application, err := entity.NewApplication(input.Name, input.Description, *user)
	if err != nil {
		return nil, err
	}

	var permissions []*entity.Permission
	for _, permission := range input.Permissions {
		p, err := entity.NewPermission(*application, permission.Resource, permission.Action, permission.CreatedBy)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}

	application.Permissions = permissions

	err = a.repository.Save(ctx, application)

	return &ApplicationOutput{
		ID:          application.ID,
		Name:        application.Name,
		Slug:        application.Slug,
		Description: application.Description,
		TeamID:      application.TeamID,
		CreatedBy:   application.CreatedBy,
	}, err
}
