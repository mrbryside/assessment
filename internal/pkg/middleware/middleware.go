package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func VerifyAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader != "November 10, 2009" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header")
		}

		return next(c)
	}
}
