package rest

import (
	"github.com/guilhermealvess/oauth-api/internal/interface/rest/handler.go"
	"github.com/guilhermealvess/oauth-api/internal/interface/rest/router"
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

	router.NewApiServerHandler(api, "api-server", handler.NewApiServerHandler()).Bind()
}
