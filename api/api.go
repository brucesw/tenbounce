package api

import (
	"tenbounce/util"

	"github.com/labstack/echo/v4"
)

type HandlerClx struct {
	repository    Repository
	signingSecret string
	nower         util.Nower
}

func NewTenbounceAPI(
	repository Repository,
	signingSecret string,
) *echo.Echo {
	var apiServer = echo.New()

	var handlerClx = HandlerClx{
		repository:    repository,
		signingSecret: signingSecret,
		nower:         util.NewTimeNower(),
	}

	uiRoutes(apiServer, handlerClx)

	setUserRoutes(apiServer, handlerClx)

	var apiGroup = apiServer.Group("/api")
	apiRoutes(apiGroup, handlerClx)

	return apiServer
}

func apiRoutes(g *echo.Group, h HandlerClx) {
	// Routes require user to be set
	g.Use(h.RequireUserMiddleware)

	// TODO(bruce): remove
	tempPostgresRoute(g)

	userRoutes(g, h)
	pointRoutes(g, h)
	pointTypeRoutes(g, h)
}
