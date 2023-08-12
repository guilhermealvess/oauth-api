package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/internal/domain"
	"github.com/guilhermealvess/oauth-api/internal/domain/entity"
)

const _team = "team"

type teamRepository struct {
	cache *redis.Client
}

func NewTeamRepository(redisClient *redis.Client) domain.TeamRepository {
	return teamRepository{
		cache: redisClient,
	}
}

func (t teamRepository) Save(ctx context.Context, team *entity.Team) error {
	key := fmt.Sprintf("%s:%s", _team, team.ID.String())
	raw, err := json.Marshal(team)
	if err != nil {
		return err
	}

	if err := t.cache.Set(ctx, key, raw, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (t teamRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Team, error) {
	key := fmt.Sprintf("%s:%s", _team, id.String())

	result := t.cache.Get(ctx, key)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var team entity.Team
	err := result.Scan(&team)
	return &team, err
}
