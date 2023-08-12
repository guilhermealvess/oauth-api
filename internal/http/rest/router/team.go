package router

import (
	"fmt"

	"github.com/guilhermealvess/oauth-api/internal/http/rest/handler.go"
	"github.com/labstack/echo/v4"
)

type TeamsRouter struct {
	resource string
	group    *echo.Group
	handler  handler.TeamsHandler
}

func NewTeamsRouter(g *echo.Group, r string, h handler.TeamsHandler) Router {
	return TeamsRouter{
		resource: r,
		group:    g,
		handler:  h,
	}
}

func (t TeamsRouter) Bind() {
	t.group.POST(fmt.Sprintf("/%s/", t.resource), t.handler.Create)
	t.group.GET(fmt.Sprintf("/%s/:id", t.resource), t.handler.Get)
}
