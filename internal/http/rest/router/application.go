package router

import (
	"fmt"

	"github.com/guilhermealvess/oauth-api/internal/http/rest/handler.go"
	"github.com/labstack/echo/v4"
)

type ApiServerRouter struct {
	resource string
	group    *echo.Group
	handler  handler.ApplicationHandler
}

func NewApiServerHandler(g *echo.Group, resource string, h handler.ApplicationHandler) Router {
	return ApiServerRouter{
		resource: resource,
		group:    g,
		handler:  h,
	}
}

func (a ApiServerRouter) Bind() {
	a.group.POST(fmt.Sprintf("/%s/", a.resource), a.handler.Create)
}
