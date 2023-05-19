package server

import (
	"digibala/models"
	"fmt"

	"github.com/labstack/echo/v4"
)

// user routes comes here

func userHandler(c echo.Context) error {
	return nil
}

func init() {
	fmt.Println(models.User{})
	e.GET("/user", userHandler)
}
