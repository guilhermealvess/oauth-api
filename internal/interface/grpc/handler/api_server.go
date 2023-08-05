package handler

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/guilhermealvess/oauth-api/internal/interface/grpc/pb"
	"github.com/guilhermealvess/oauth-api/internal/repository"
	"github.com/guilhermealvess/oauth-api/internal/usecase"
)

type ApiServerHandler struct {
	pb.UnimplementedApiServerServer
	usecase usecase.ApiUsecase
}

func NewApiServerHandler() *ApiServerHandler {
	return &ApiServerHandler{
		usecase: usecase.NewApiUse(
			repository.NewApiServerRepository(redis.NewClient(&redis.Options{})),
		),
	}
}

func (a *ApiServerHandler) Create(ctx context.Context, req *pb.CreateApiServerRequest) (*pb.CreateApiServerResponse, error) {
	in := usecase.CreateApiServerInput{
		Name:        req.Name,
		Description: req.Desc,
		CreatedBy:   req.CreatedBy,
	}

	out, err := a.usecase.CreateApiServer(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateApiServerResponse{
		ID:        out.Name,
		Desc:      &out.Description,
		Slug:      out.Slug,
		IsActive:  out.IsActive,
		CreatedBy: out.CreatedBy,
		UpdatedAt: out.UpdatedAt.Local().String(),
	}, nil
}
