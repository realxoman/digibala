package main

import (
	"faq/handler"
	"github.com/labstack/echo/v4"
	"log"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func main() {
	server.FaqRoutes(e)
	log.Fatal(e.Start(":8080"))
}
