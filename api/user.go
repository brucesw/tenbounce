package api

import (
	_ "embed"
	"fmt"
	"net/http"
	"tenbounce/model"
	"time"

	"encoding/json"

	"github.com/labstack/echo/v4"
)

const UserIDCookieName string = "UserID"

// TODO(bruce): XXX
var hardcodedUsers []userWithSecretURL

type userWithSecretURL struct {
	model.User

	SecretURL string `json:"secretURL"`
}

func userRoutes(g *echo.Group) {
	var userRoutes = g.Group("/users")

	// TODO(bruce): XXX
	var setUserRoutes = userRoutes.Group("/set_user")

	for _, hardcodedUser := range hardcodedUsers {
		setUserRoutes.GET("/"+hardcodedUser.SecretURL, func(c echo.Context) error {
			var cookie = new(http.Cookie)
			cookie.Name = UserIDCookieName
			cookie.Value = hardcodedUser.ID
			cookie.Path = "/"
			// TODO(bruce): User nower, determine expiration time
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.SetCookie(cookie)

			return c.JSON(http.StatusOK, "logged in as "+hardcodedUser.Name)
		})
	}

}

//go:embed user_secrets.json
var hardcodedUsers_bytes []byte

func init() {
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}
}
