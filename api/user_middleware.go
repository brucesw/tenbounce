package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequireUserMiddleware ensures a valid cookie is present, erroring otherwise.
// Adds the user ID to the request context for downstream consumption.
func (h HandlerClx) RequireUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if userIDCookie, err := c.Cookie(CookieName_UserID); err != nil {
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
