package api

import (
	_ "embed"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed user_secrets.json
var hardcodedUsers_bytes []byte

// TODO(bruce): XXX
var hardcodedUsers []userWithSecretURL

var APIServer = echo.New()
var apiGroup = APIServer.Group("/api")

func apiRoutes(g *echo.Group) {
	pointRoutes(g)
	userRoutes(g)

	tempPostgresRoute(g)
}

func init() {
	// TODO(bruce): XXX needs to run before routes are registered so users exist
	// api.go init() runs before user.go init()
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}

	apiRoutes(apiGroup)

	// TODO(bruce): UI routes
	APIServer.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, homepageHTML)
	})
}
