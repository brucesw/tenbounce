package api

import (
	"net/http"
	"tenbounce/util"

	"github.com/labstack/echo/v4"
)

type HandlerClx struct {
	repository    Repository
	signingSecret string
	nower         util.Nower
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
		nower:         util.NewTimeNower(),
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
