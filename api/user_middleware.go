package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO(bruce): implement
// TODO(bruce): document
func (h HandlerClx) SetUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())

		if userIDCookie, err := c.Cookie(userIDCookieName); err != nil {
			// TODO(bruce): returns error
			// TODO(bruce): redirect to login page
			return c.JSON(http.StatusUnauthorized, "get user id cookie")
		} else if userID, err := userID_FromCookieValue(userIDCookie.Value, h.signingSecret); err != nil {
			return c.JSON(http.StatusUnauthorized, "user id from cookie value")
		} else {
			var newCtx = contextWithUserID(c.Request().Context(), userID)
			var newRequestContext = c.Request().WithContext(newCtx)
			c.SetRequest(newRequestContext)

			if err := next(c); err != nil {
				c.Error(err)
			}
		}

		return nil
	}
}
