package api

import (
	"github.com/labstack/echo/v4"
)

var APIServer = echo.New()
var apiGroup = APIServer.Group("/api")

func apiRoutes(g *echo.Group) {
	pointRoutes(g)
}

func init() {
	apiRoutes(apiGroup)
}
