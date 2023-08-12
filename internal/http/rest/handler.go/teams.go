package handler

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/config"
	"github.com/guilhermealvess/oauth-api/internal/repository"
	"github.com/guilhermealvess/oauth-api/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TeamsHandler struct {
	usecase usecase.TeamUseCase
}

func NewTeamsHandler() *TeamsHandler {
	return &TeamsHandler{
		usecase: usecase.NewTeamUseCase(
			repository.NewTeamRepository(redis.NewClient(&redis.Options{Addr: config.Config.RedisAddress, DB: 0})),
		),
	}
}

func (t *TeamsHandler) Create(c echo.Context) error {
	var payload struct {
		Name string `json:"name"`
		Desc string `json:"description"`
	}

	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	output, err := t.usecase.Create(c.Request().Context(), payload.Name, payload.Desc)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &output)
}

func (t *TeamsHandler) Get(c echo.Context) error {
	teamID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	out, err := t.usecase.FindById(c.Request().Context(), teamID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &out)
}
