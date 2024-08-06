package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func pointTypeRoutes(g *echo.Group, h HandlerClx) {
	var pointTypeRoutes = g.Group("/point_types")

	pointTypeRoutes.GET("", h.listPointTypes)

}

// TODO(bruce): document
// TODO(bruce): responses
func (h HandlerClx) listPointTypes(c echo.Context) error {

	// TODO(bruce): confirm valid user

	if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get point types from db")
	} else {
		return c.JSON(http.StatusOK, pointTypes)
	}
}
