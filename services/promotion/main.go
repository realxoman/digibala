package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	database "promotion/database"
	server "promotion/server"
)

func main() {
	// initialize the database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Auto-migrate
	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true
	bs, _ := os.ReadFile("server/promotion.txt")
	fmt.Println(string(bs))

	promotionService := server.NewPromotionService(db)
	server.RegisterRoutes(e, promotionService)

	log.Fatal(e.Start(":8787"))
}
