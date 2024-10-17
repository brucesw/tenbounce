package api

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed ui/homepage.html
var homepageHTML string

//go:embed ui/unauthorized.html
var unauthorizedHTML string

func uiRoutes(e *echo.Echo, h HandlerClx) {
	e.GET("", func(c echo.Context) error {
		if userIDCookie, err := c.Cookie(CookieName_UserID); err != nil {
			return c.HTML(http.StatusUnauthorized, unauthorizedHTML)
		} else if userID, err := userID_FromCookieValue(userIDCookie.Value, h.signingSecret); err != nil {
			return c.HTML(http.StatusUnauthorized, unauthorizedHTML)
		} else if _, err := h.repository.GetUser(userID); err != nil {
			return c.HTML(http.StatusUnauthorized, unauthorizedHTML)
		}

		return c.HTML(http.StatusOK, homepageHTML)
	})
}
