package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/internal/domain"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type userRepository struct {
	cache *redis.Client
}

func NewUserRpository(redisClient *redis.Client) domain.UserRepository {
	return &userRepository{
		cache: redisClient,
	}
}

func (u *userRepository) Save(ctx context.Context, user *entity.User) error {
	keyID := fmt.Sprintf("user:%s", user.ID.String())
	keyEmail := fmt.Sprintf("user:%s", user.Email)

	if err := u.cache.Set(ctx, keyID, user, 0).Err(); err != nil {
		return err
	}

	return u.cache.Set(ctx, keyEmail, user, 0).Err()
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	key := fmt.Sprintf("user:%s", email)
	result := u.cache.Get(ctx, key)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var user entity.User
	if err := result.Scan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindByID(context.Context, uuid.UUID) (*entity.User, error) {
	return nil, nil
}
