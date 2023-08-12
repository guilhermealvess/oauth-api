package rest

import (
	"github.com/guilhermealvess/oauth-api/internal/http/rest/handler.go"
	"github.com/guilhermealvess/oauth-api/internal/http/rest/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewApp() *echo.Echo {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	bindRouters(app)

	return app
}

func bindRouters(app *echo.Echo) {
	api := app.Group("/api")

	router.NewApiServerHandler(api, "application", handler.NewApplicationHandler()).Bind()
	router.NewTeamsRouter(api, "teams", *handler.NewTeamsHandler()).Bind()
}
