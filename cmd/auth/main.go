package main

import (
	"auth/routes"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load("cmd/auth/.env")
	e := echo.New()
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}
