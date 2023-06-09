package main

import (
	"admin/handler"
	"log"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func main() {
	handler.AdminRoutes(e)
	log.Fatal(e.Start(":8081"))
}
