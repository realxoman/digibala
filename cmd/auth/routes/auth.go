package routes

import (
	"net/http"

	"auth/utils"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	authGroup := e.Group("/api/auth")
	authGroup.POST("/login", Login)
	// for refresh token and logout we must use databases for revoking tokens
	authGroup.POST("/refresh", RefreshToken)

}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if password != "admin" && username != "admin" {
		return echo.ErrUnauthorized
	}
	accessToken, refreshToken, err := utils.CreateTokens("admin")
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c echo.Context) error {
	ok, claims, err := utils.IsValidToken(c.Request().Header.Get("Authorization"))
	if err != nil || !ok {
		return echo.ErrUnauthorized
	}
	if claims.TokenType != "refresh_type" {
		return echo.ErrBadRequest
	}
	accessToken, refreshToken, tokenErr := utils.CreateTokens(claims.Username)
	if tokenErr != nil {
		return echo.ErrBadGateway.Internal
	}
	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
