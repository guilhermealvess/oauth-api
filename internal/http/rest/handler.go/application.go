package handler

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/guilhermealvess/oauth-api/config"
	"github.com/guilhermealvess/oauth-api/internal/repository"
	"github.com/guilhermealvess/oauth-api/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ApplicationHandler struct {
	usecase usecase.ApplicationUseCase
}

func NewApplicationHandler() ApplicationHandler {
	cli := &redis.Options{
		Addr: config.Config.RedisAddress,
		DB:   0,
	}

	return ApplicationHandler{
		usecase: usecase.NewApplicationUseCase(repository.NewApiServerRepository(redis.NewClient(cli)),
			repository.NewUserRpository(redis.NewClient(cli))),
	}
}

func (a ApplicationHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var request usecase.CreateApplication
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	output, err := a.usecase.CreateApplication(ctx, request)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &output)
}
