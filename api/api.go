package api

import (
	_ "embed"
	"tenbounce/repository"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed user_secrets.json
var hardcodedUsers_bytes []byte

// TODO(bruce): XXX
var hardcodedUsers []userWithSecretURL

type HandlerClx struct {
	repository Repository
}

func apiRoutes(g *echo.Group, h HandlerClx) {
	userRoutes(g)
	tempPostgresRoute(g)

	// Routes require user to be set
	g.Use(SetUserMiddleware)
	pointRoutes(g, h)
	pointTypeRoutes(g, h)
}

func init() {
	// TODO(bruce): XXX needs to run before routes are registered so users exist
	// api.go init() runs before user.go init()
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}
}


func NewTenbounceAPI() *echo.Echo {
	var APIServer = echo.New()
	var apiGroup = APIServer.Group("/api")

	var handlerClx = HandlerClx{
		repository: repository.NewInMemoryRepository(),
	}

	apiRoutes(apiGroup, handlerClx)

	// TODO(bruce): UI routes
	APIServer.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, homepageHTML)
	})

	return APIServer
}
