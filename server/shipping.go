package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func shippingRoutes(e *echo.Echo) {
	e.GET("/shipping", listShippingHandler)
	e.GET("/shipping/:id", findShippingHandler)
	e.POST("/shipping", createShippingHandler)
	e.PUT("/shipping", updateShippingHandler)
	e.DELETE("/shipping/:id", deleteShippingHandler)
}

func updateShippingHandler(c echo.Context) error {
	shipping := &models.Address{}
	if err := c.Bind(shipping); err != nil {
		return err
	}
	fmt.Println("Updating shipping id:", shipping.ID)
	return c.JSON(http.StatusOK, shipping)
}

func deleteShippingHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting shipping id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createShippingHandler(c echo.Context) error {
	shipping := &models.Address{}
	err := c.Bind(shipping)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, shipping)
}

func listShippingHandler(c echo.Context) error {
	var shippings []models.Shipping
	return c.JSON(http.StatusOK, shippings)
}

func findShippingHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	shipping := &models.Shipping{ID: id}
	return c.JSON(http.StatusOK, shipping)
}
