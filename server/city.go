package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func cityRoutes(e *echo.Echo) {
	e.GET("/city", listCityHandler)
	e.POST("/city", createCityHandler)
	e.GET("/City/:id", findCityHandler)
	e.DELETE("/City/:id", deleteCityHandler)
	e.PUT("/City", updateCityHandler)
}

func updateCityHandler(c echo.Context) error {
	City := &models.City{}
	if err := c.Bind(City); err != nil {
		return err
	}
	fmt.Println("Updating City id:", City.ID)
	return c.JSON(http.StatusOK, City)
}

func deleteCityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting City id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createCityHandler(c echo.Context) error {
	City := &models.City{}
	if err := c.Bind(City); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, City)
}

func listCityHandler(c echo.Context) error {
	Cityes := []models.City{}
	return c.JSON(http.StatusOK, Cityes)
}

func findCityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	City := &models.City{ID: id}
	return c.JSON(http.StatusOK, City)
}
