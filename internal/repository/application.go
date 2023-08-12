package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/guilhermealvess/oauth-api/internal/domain"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

type applicationRepository struct {
	cacheClient *redis.Client
}

func NewApiServerRepository(r *redis.Client) domain.ApplicationRepository {
	return applicationRepository{r}
}

func (a applicationRepository) Save(ctx context.Context, apiServer *entity.Application) error {
	key := apiServer.ID.String()
	value, err := json.Marshal(*apiServer)
	if err != nil {
		return err
	}
	if err = a.cacheClient.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}
