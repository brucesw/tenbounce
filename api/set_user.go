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

//go:embed user_secrets.json
var hardcodedUsers_bytes []byte

var hardcodedUsers []userWithSecretURL

// Until user creation and auth is in place, leverage hardcoded
// list of users, each with a secret URL used for login
type userWithSecretURL struct {
	model.User

	SecretURL string `json:"secretURL"`
}

// register one "secret" endpoint for each hardcoded user
// endpoint sets a signed cookie indicating user is logged in
func setUserRoutes(e *echo.Echo, h HandlerClx) {
	var setUserRoutes = e.Group("/set_user")

	for _, hardcodedUser := range hardcodedUsers {
		setUserRoutes.GET("/"+hardcodedUser.SecretURL, func(c echo.Context) error {
			var cookie = new(http.Cookie)
			cookie.Name = userIDCookieName

			cookieValue, err := userID_ToCookieValue(hardcodedUser.ID, h.signingSecret)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Errorf("user id to cookie value: %w", err))
			}
			cookie.Value = cookieValue

			cookie.Path = "/"
			// 7 Days till expiration
			cookie.Expires = h.nower.Now().Add(7 * 24 * time.Hour)
			c.SetCookie(cookie)

			return c.JSON(http.StatusOK, "logged in as "+hardcodedUser.Name)
		})
	}

}

func init() {
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}
}
