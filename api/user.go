package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func userRoutes(g *echo.Group, h HandlerClx) {
	var userRoutes = g.Group("/users")

	userRoutes.GET("/me", h.getCurrentUser)
	userRoutes.GET("", h.listUsers)
}

func (h HandlerClx) getCurrentUser(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if user, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func (h HandlerClx) listUsers(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else if users, err := h.repository.ListUsers(); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get users: %w", err))
	} else {
		return c.JSON(http.StatusOK, users)
	}
}
