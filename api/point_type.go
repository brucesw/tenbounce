package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func pointTypeRoutes(g *echo.Group) {
	var pointTypeRoutes = g.Group("/point_types")

	pointTypeRoutes.GET("", listPointTypes)

}

// TODO(bruce): document
// TODO(bruce): responses
func listPointTypes(c echo.Context) error {
	if pointTypes, err := GetPointTypesFromDB(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get point types from db")
	} else {
		return c.JSON(http.StatusOK, pointTypes)
	}

}
