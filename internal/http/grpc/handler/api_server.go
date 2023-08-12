package handler

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/guilhermealvess/oauth-api/internal/repository"
	"github.com/guilhermealvess/oauth-api/internal/usecase"
	"github.com/guilhermealvess/oauth-api/pkg/grpc/pb"
)

type ApplicationHandler struct {
	pb.UnimplementedApplicationServer
	usecase usecase.ApplicationUseCase
}

func NewApiServerHandler() *ApplicationHandler {
	redisClient := redis.NewClient(&redis.Options{})
	return &ApplicationHandler{
		usecase: usecase.NewApplicationUseCase(
			repository.NewApiServerRepository(redisClient),
			repository.NewUserRpository(redisClient),
		),
	}
}

func (a *ApplicationHandler) Create(ctx context.Context, req *pb.CreateApplicationRequest) (*pb.CreateApplicationResponse, error) {
	input := usecase.CreateApplication{
		UserEmail:   "email@todo.com",
		Name:        req.Name,
		Description: req.Desc,
	}

	output, err := a.usecase.CreateApplication(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateApplicationResponse{
		ID:          output.ID.String(),
		Name:        output.Name,
		Description: &output.Description,
		CreatedBy:   output.CreatedBy,
	}, nil
}
