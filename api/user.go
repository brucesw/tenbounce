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
	// TODO(bruce): helper function to get user id off context
	var userID = ctx.Value(ctxKey_UserID("userID"))
	var userIDString string
	var ok bool
	if userIDString, ok = userID.(string); !ok {
		return c.JSON(http.StatusInternalServerError, "user id not string type")
	}

	if user, err := h.repository.GetUser(userIDString); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else {
		return c.JSON(http.StatusOK, user)
	}
}
