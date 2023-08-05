package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
	"github.com/guilhermealvess/oauth-api/internal/domain/repository"
)

type apiServerRepository struct {
	cacheClient *redis.Client
}

func NewApiServerRepository(r *redis.Client) repository.ApiServerRepository {
	return apiServerRepository{r}
}

func (a apiServerRepository) Save(ctx context.Context, apiServer *entity.ApiServer) error {
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
