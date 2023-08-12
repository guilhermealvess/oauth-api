package usecase

import (
	"context"

	"github.com/guilhermealvess/oauth-api/internal/domain"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type UserUseCase interface {
	Create(ctx context.Context, input CreateUserInput) (*UserOutput, error)
	UserSignIn(ctx context.Context, email, password string) (string, error)
}

type userUse struct {
	repository domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) UserUseCase {
	return &userUse{repo}
}

func (u *userUse) Create(ctx context.Context, input CreateUserInput) (*UserOutput, error) {
	user, err := entity.NewUser(input.TeamID, input.Name, input.Occupation, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	if err := user.Ok(); err != nil {
		return nil, err
	}

	if err := u.repository.Save(ctx, user); err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:         user.ID,
		TeamID:     user.TeamID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Occupation,
	}, nil
}

func (u *userUse) UserSignIn(ctx context.Context, email, password string) (string, error) {
	user, err := u.repository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return user.SignIn(password)
}
