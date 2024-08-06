package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO(bruce): implement
// TODO(bruce): document
func SetUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())

		if userIDCookie, err := c.Cookie(UserIDCookieName); err != nil {
			// TODO(bruce): returns error
			// TODO(bruce): redirect to login page
			return c.JSON(http.StatusUnauthorized, "get user id cookie")
		} else if true == false {
			// TODO(bruce): implement cookie hashing with secret validation
		} else {
			var newCtx = contextWithUserID(c.Request().Context(), userIDCookie.Value)
			var newRequestContext = c.Request().WithContext(newCtx)
			c.SetRequest(newRequestContext)

			if err := next(c); err != nil {
				c.Error(err)
			}
		}

		return nil
	}
}
