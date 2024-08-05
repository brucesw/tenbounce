package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// TODO(bruce): implement
// TODO(bruce): document
func SetUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
