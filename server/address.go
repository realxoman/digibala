package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddressRoutes(e *echo.Echo) {
	e.GET("/address", listAddressHandler)
	e.POST("/address", createAddressHandler)
	e.GET("/address/:id", findAddressHandler)
	e.DELETE("/address/:id", deleteAddressHandler)
	e.PUT("/address", updateAddressHandler)
}

func updateAddressHandler(c echo.Context) error {
	address := &models.Address{}
	if err := c.Bind(address); err != nil {
		return err
	}
	fmt.Println("Updating address id:", address.ID)
	return c.JSON(http.StatusOK, address)
}

func deleteAddressHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting address id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createAddressHandler(c echo.Context) error {
	address := &models.Address{}
	if err := c.Bind(address); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, address)
}

func listAddressHandler(c echo.Context) error {
	var addresses []models.Address
	return c.JSON(http.StatusOK, addresses)
}

func findAddressHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	address := &models.Address{ID: id}
	return c.JSON(http.StatusOK, address)
}
