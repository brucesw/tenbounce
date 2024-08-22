package api

import (
	"tenbounce/util"
	"time"

	"github.com/labstack/echo/v4"
)

type HandlerClx struct {
	repository    Repository
	signingSecret string
	startupTime   time.Time
	nower         util.Nower
}

func newHandlerClx(
	repository Repository,
	signingSecret string,
) HandlerClx {
	var nower = util.NewTimeNower()
	var startupTime = nower.Now()

	return HandlerClx{
		repository:    repository,
		signingSecret: signingSecret,
		startupTime:   startupTime,
		nower:         nower,
	}

}

func NewTenbounceAPI(
	repository Repository,
	signingSecret string,
) *echo.Echo {
	var apiServer = echo.New()

	var handlerClx = newHandlerClx(repository, signingSecret)

	uiRoutes(apiServer, handlerClx)

	setUserRoutes(apiServer, handlerClx)

	healthRoute(apiServer, handlerClx)

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
