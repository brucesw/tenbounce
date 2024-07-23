package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func pointRoutes(g *echo.Group) {
	var pointRoutes = g.Group("/points")

	pointRoutes.POST("", createPoint)
	pointRoutes.GET("", listPoints)
}

func createPoint(c echo.Context) error {

	return c.JSON(http.StatusOK, "create")
}

func listPoints(c echo.Context) error {

	return c.JSON(http.StatusOK, "list")
}
