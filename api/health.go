package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthCheck struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func healthRoute(e *echo.Echo, h HandlerClx) {
	e.GET("/health", func(c echo.Context) error {

		var healthCheck = HealthCheck{
			Status:    http.StatusOK,
			Timestamp: h.nower.Now(),
		}

		return c.JSON(http.StatusOK, healthCheck)
	})
}
