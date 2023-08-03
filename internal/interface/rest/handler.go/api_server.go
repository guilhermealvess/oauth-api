package handler

import (
	"net/http"

	"github.com/guilhermealvess/oauth-api/internal/repository"
	"github.com/guilhermealvess/oauth-api/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ApiServerHandler struct {
	usecase usecase.ApiUsecase
}

func NewApiServerHandler() ApiServerHandler {
	return ApiServerHandler{
		usecase: usecase.NewApiUse(repository.NewApiServerRepository()),
	}
}

func (a ApiServerHandler) Create(c echo.Context) error {
	var request usecase.CreateApiServerInput
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	output, err := a.usecase.CreateApiServer(request)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &output)
}
