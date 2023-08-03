package router

import (
	"fmt"

	"github.com/guilhermealvess/oauth-api/internal/interface/rest/handler.go"
	"github.com/labstack/echo/v4"
)

type ApiServerRouter struct {
	resource string
	group    *echo.Group
	handler  handler.ApiServerHandler
}

func NewApiServerHandler(g *echo.Group, resource string, h handler.ApiServerHandler) ApiServerRouter {
	return ApiServerRouter{
		resource: resource,
		group:    g,
		handler:  h,
	}
}

func (a ApiServerRouter) Bind() {
	a.group.POST(fmt.Sprintf("/%s/", a.resource), a.handler.Create)
}
