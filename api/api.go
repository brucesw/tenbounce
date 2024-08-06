package api

import (
	_ "embed"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

//go:embed user_secrets.json
var hardcodedUsers_bytes []byte

// TODO(bruce): XXX
var hardcodedUsers []userWithSecretURL

type HandlerClx struct {
	repository    Repository
	signingSecret string
}

func apiRoutes(g *echo.Group, h HandlerClx) {
	// TODO(bruce): remove
	tempPostgresRoute(g)

	setUserRoutes(g, h)

	// Routes require user to be set
	g.Use(h.SetUserMiddleware)
	userRoutes(g, h)
	pointRoutes(g, h)
	pointTypeRoutes(g, h)

}

func NewTenbounceAPI(
	repository Repository,
	signingSecret string,
) *echo.Echo {
	var apiServer = echo.New()
	var apiGroup = apiServer.Group("/api")

	var handlerClx = HandlerClx{
		repository:    repository,
		signingSecret: signingSecret,
	}

	apiRoutes(apiGroup, handlerClx)

	apiServer.GET("", func(c echo.Context) error {
		if userIDCookie, err := c.Cookie(userIDCookieName); err != nil {
			return c.HTML(http.StatusUnauthorized, unauthorizedHTML)
		} else if userID, err := userID_FromCookieValue(userIDCookie.Value, signingSecret); err != nil {
			return c.HTML(http.StatusUnauthorized, unauthorizedHTML)
		} else if _, err := repository.GetUser(userID); err != nil {
			return c.HTML(http.StatusUnauthorized, unauthorizedHTML)
		}

		return c.HTML(http.StatusOK, homepageHTML)
	})

	return apiServer
}

func init() {
	// TODO(bruce): XXX needs to run before routes are registered so users exist
	// api.go init() runs before user.go init()
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}
}
