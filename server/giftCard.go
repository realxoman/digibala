package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func giftCardRoutes(e *echo.Echo) {
	e.GET("/giftCards", listgiftCardsHandler)
	e.POST("/giftCard", creategiftCardsHandler)
	e.GET("/giftCards/:id", findgiftCardsHandler)
	e.DELETE("/giftCards/:id", deletegiftCardsHandler)
	e.PUT("/giftCards", updategiftCardsHandler)
}

func updategiftCardsHandler(c echo.Context) error {
	giftCard := &models.GiftCard{}
	if err := c.Bind(giftCard); err != nil {
		return err
	}
	fmt.Println("Updating giftCards id:", giftCard.ID)
	return c.JSON(http.StatusOK, giftCard)
}

func deletegiftCardsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting giftCards id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func creategiftCardsHandler(c echo.Context) error {
	giftCard := &models.GiftCard{}
	if err := c.Bind(giftCard); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, giftCard)
}

func listgiftCardsHandler(c echo.Context) error {
	giftCardses := []models.GiftCard{}
	return c.JSON(http.StatusOK, giftCardses)
}

func findgiftCardsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	giftCards := &models.GiftCard{ID: id}
	return c.JSON(http.StatusOK, giftCards)
}
