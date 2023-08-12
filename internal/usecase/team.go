package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/internal/domain"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type TeamUseCase interface {
	Create(ctx context.Context, name, desc string) (*TeamOutput, error)
	FindById(ctx context.Context, id uuid.UUID) (*TeamOutput, error)
}

type teamUse struct {
	repository domain.TeamRepository
}

func NewTeamUseCase(repo domain.TeamRepository) TeamUseCase {
	return &teamUse{
		repository: repo,
	}
}

func (t *teamUse) Create(ctx context.Context, name, desc string) (*TeamOutput, error) {
	team, err := entity.NewTeam(name, desc)
	if err != nil {
		return nil, err
	}

	if err := t.repository.Save(ctx, team); err != nil {
		return nil, err
	}

	return &TeamOutput{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
	}, nil
}

func (t *teamUse) FindById(ctx context.Context, id uuid.UUID) (*TeamOutput, error) {
	team, err := t.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := team.Ok(); err != nil {
		return nil, err
	}

	if err := t.repository.Save(ctx, team); err != nil {
		return nil, err
	}

	return &TeamOutput{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
	}, nil
}
