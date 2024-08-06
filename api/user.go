package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func userRoutes(g *echo.Group, h HandlerClx) {
	var userRoutes = g.Group("/users")

	userRoutes.GET("", h.getUser)
}

func (h HandlerClx) getUser(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if user, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else {
		return c.JSON(http.StatusOK, user)
	}
}
