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

type HealthCheck struct {
	CommitSha string    `json:"commitSha"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

// TODO(bruce): https://bdemirpolat.medium.com/golang-compiler-directives-dc61820add40
func healthRoute(e *echo.Echo, h HandlerClx) {
	e.GET("/health", func(c echo.Context) error {

		var uptime = h.nower.Now().Sub(h.startupTime)

		var healthCheck = HealthCheck{
			Status:    http.StatusOK,
			Timestamp: h.nower.Now(),
			CommitSha: commitSha,
			Uptime:    uptime.String(),
		}

		return c.JSON(http.StatusOK, healthCheck)
	})
}
