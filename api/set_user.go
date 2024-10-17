package api

import (
	_ "embed"

	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// register one "secret" endpoint for each hardcoded user
// endpoint sets a signed cookie indicating user is logged in
func setUserRoutes(e *echo.Echo, h HandlerClx) {
	var setUserRoutes = e.Group("/set_user")

	for _, hardcodedUser := range h.tempHardcodedUsers {
		setUserRoutes.GET("/"+hardcodedUser.SecretURL, func(c echo.Context) error {
			var cookie = new(http.Cookie)
			cookie.Name = CookieName_UserID

			cookieValue, err := userID_ToCookieValue(hardcodedUser.ID, h.signingSecret)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Errorf("user id to cookie value: %w", err))
			}
			cookie.Value = cookieValue

			cookie.Path = "/"
			// 7 Days till expiration
			cookie.Expires = h.nower.Now().Add(7 * 24 * time.Hour)
			c.SetCookie(cookie)

			return c.Redirect(http.StatusFound, "/")
		})
	}

}
