package cmd

import (
	"digibala/server"
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	// Middleware
	server.AddressRoutes(e)
	e.Start(":8080")

}
