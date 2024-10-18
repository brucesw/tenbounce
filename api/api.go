package api

import (
	"tenbounce/model"
	"tenbounce/util"
	"time"

	"github.com/labstack/echo/v4"
)

// Until user creation and auth is in place, leverage hardcoded
// list of users, each with a secret URL used for login
type UserWithSecretURL struct {
	model.User

	SecretURL string `json:"secretURL"`
}

type HandlerClx struct {
	repository         Repository
	signingSecret      string
	startupTime        time.Time
	tempHardcodedUsers []UserWithSecretURL
	nower              util.Nower
	openAIAPIKey       string
}

func newHandlerClx(
	repository Repository,
	signingSecret string,
	tempHardcodedUsers []UserWithSecretURL,
	openAIAPIKey string,
) HandlerClx {
	var nower = util.NewTimeNower()
	var startupTime = nower.Now()

	return HandlerClx{
		repository:         repository,
		signingSecret:      signingSecret,
		startupTime:        startupTime,
		tempHardcodedUsers: tempHardcodedUsers,
		nower:              nower,
		openAIAPIKey:       openAIAPIKey,
	}
}

func NewTenbounceAPI(
	repository Repository,
	signingSecret string,
	tempHardcodedUsers []UserWithSecretURL,
	openAIAPIKey string,
) *echo.Echo {
	var apiServer = echo.New()

	var handlerClx = newHandlerClx(
		repository,
		signingSecret,
		tempHardcodedUsers,
		openAIAPIKey,
	)

	uiRoutes(apiServer, handlerClx)

	setUserRoutes(apiServer, handlerClx)

	healthRoute(apiServer, handlerClx)

	var apiGroup = apiServer.Group("/api")
	apiRoutes(apiGroup, handlerClx)

	var statsGroup = apiServer.Group("/stats")
	statsRoutes(statsGroup, handlerClx)

	return apiServer
}

func apiRoutes(g *echo.Group, h HandlerClx) {
	// Routes require user to be set
	g.Use(h.RequireUserMiddleware)

	userRoutes(g, h)
	pointRoutes(g, h)
	pointTypeRoutes(g, h)
}

func statsRoutes(g *echo.Group, h HandlerClx) {
	// Routes require user to be set
	g.Use(h.RequireUserMiddleware)

	statsRoutez(g, h)
}
