package middelwares

import (
	"auth/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ok, claims, err := utils.IsValidToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if ok {
			c.Request().Header.Set("username", claims.Username)
			return next(c)
		}
		return echo.ErrUnauthorized

	}
}
