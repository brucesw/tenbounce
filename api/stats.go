package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func statsRoutez(g *echo.Group, h HandlerClx) {
	g.GET("/summary", h.getStatsSummary)
}

func (h HandlerClx) getStatsSummary(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
		// TODO(bruce): confirm creator user has permission to create points for user
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if statsSummary, err := h.repository.GetStatsSummary(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get stats summary")
	} else {
		return c.JSON(http.StatusOK, statsSummary)
	}
}
