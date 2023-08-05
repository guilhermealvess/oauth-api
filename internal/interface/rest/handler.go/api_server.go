package handler

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/guilhermealvess/oauth-api/internal/repository"
	"github.com/guilhermealvess/oauth-api/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ApiServerHandler struct {
	usecase usecase.ApiUsecase
}

func NewApiServerHandler() ApiServerHandler {
	return ApiServerHandler{
		usecase: usecase.NewApiUse(repository.NewApiServerRepository(redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}))),
	}
}

func (a ApiServerHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var request usecase.CreateApiServerInput
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	output, err := a.usecase.CreateApiServer(ctx, request)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &output)
}
