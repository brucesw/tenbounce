package api

import "github.com/labstack/echo/v4"
import "net/http"

var APIServer = echo.New()

func init() {
	// TODO(bruce): Separate routes
	APIServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
