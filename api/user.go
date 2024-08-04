package api

import (
	"fmt"
	"net/http"
	"tenbounce/model"
	"time"

	"encoding/json"

	"github.com/labstack/echo/v4"
)

const UserIDCookieName string = "TENBOUNCE_USER_ID"

// Until user creation and auth is in place, leverage hardcoded
// list of users, each with a secret URL used for login
type userWithSecretURL struct {
	model.User

	SecretURL string `json:"secretURL"`
}

func userRoutes(g *echo.Group) {
	var userRoutes = g.Group("/users")

	// TODO(bruce): XXX one route per user for login
	var setUserRoutes = userRoutes.Group("/set_user")

	for _, hardcodedUser := range hardcodedUsers {
		setUserRoutes.GET("/"+hardcodedUser.SecretURL, func(c echo.Context) error {
			var cookie = new(http.Cookie)
			cookie.Name = UserIDCookieName
			// TODO(bruce): cookie value needs to have some sort of hash in it for security + verification middleware needed
			cookie.Value = hardcodedUser.ID
			cookie.Path = "/"
			// TODO(bruce): User nower, determine expiration time
			cookie.Expires = time.Now().Add(24 * time.Hour)
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
