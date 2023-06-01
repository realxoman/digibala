package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func stockRoutes(e *echo.Echo) {
	e.GET("/stock", listStockHandler)
	e.POST("/stock", createStockHandler)
	e.GET("/stock:id", findStockHandler)
	e.DELETE("/stock:id", deleteStockHandler)
	e.PUT("/stock:id", updateStockHandler)
}

func listStockHandler(c echo.Context) error {
	var stocks []models.Stock
	return c.JSON(http.StatusOK, stocks)
}
func findStockHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// quantity := &models.Stock{ProductID: id}
	return c.JSON(http.StatusOK, id)
}
func createStockHandler(c echo.Context) error {
	stock := &models.Stock{}
	if err := c.Bind(stock); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, stock)
}
func updateStockHandler(c echo.Context) error {
	stock := &models.Stock{}
	id, _ := strconv.Atoi(c.Param("id"))
	quantity := &models.Stock{ProductID: id}
	if err := c.Bind(stock); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Updating stock id:", id)
	return c.JSON(http.StatusOK, quantity)
}
func deleteStockHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting stock id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}
