package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var APIServer = echo.New()
var apiGroup = APIServer.Group("/api")

func apiRoutes(g *echo.Group) {
	pointRoutes(g)
}

func init() {
	apiRoutes(apiGroup)

	APIServer.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, homepageHTML)
	})
}
