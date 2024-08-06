package api

import (
	"fmt"
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
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get point types from db")
	} else {
		return c.JSON(http.StatusOK, pointTypes)
	}
}
