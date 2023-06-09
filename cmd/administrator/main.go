package main

import (
	"admin/server"
	"log"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func main() {
	server.AdminRoutes(e)
	log.Fatal(e.Start(":8081"))
}
