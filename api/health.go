package api

import (
	"net/http"
	"time"

	_ "embed"

	"github.com/labstack/echo/v4"
)

//go:generate ./generate_commit_sha.sh

//go:embed commit_sha.txt
var commitSha string

//go:embed clean.txt
var clean string

type HealthCheck struct {
	Clean     string    `json:"clean"`
	CommitSha string    `json:"commitSha"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

func healthRoute(e *echo.Echo, h HandlerClx) {
	e.GET("/health", func(c echo.Context) error {
		var uptime = h.nower.Now().Sub(h.startupTime)

		var healthCheck = HealthCheck{
			Clean:     clean,
			CommitSha: commitSha,
			Status:    http.StatusOK,
			Timestamp: h.nower.Now(),
			Uptime:    uptime.String(),
		}

		return c.JSON(http.StatusOK, healthCheck)
	})
}
