package server

import (
	"digibala/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func listCurrencyHandler(c echo.Context) error {
	var currencies []*models.Currency
	return c.JSON(http.StatusOK, currencies)
}

func getCurrencyByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	currency := models.Currency{
		ID: id,
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	return c.JSON(http.StatusOK, currency)
}

func createCurrencyHandler(c echo.Context) error {
	currency := new(models.Currency)

	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency data")
	}

	return c.JSON(http.StatusCreated, currency)
}

func updateCurrencyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	currency := new(models.Currency)
	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency data")
	}

	currency.ID = id

	return c.JSON(http.StatusOK, currency)
}

func deleteCurrencyHandler(c echo.Context) error {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	return c.NoContent(http.StatusOK)
}

func currencyRoutes(e *echo.Echo) {
	e.GET("/api/currency", listCurrencyHandler)
	e.GET("/api/currency/:id", getCurrencyByIDHandler)
	e.POST("/api/currency", createCurrencyHandler)
	e.PUT("/api/currency/:id", updateCurrencyHandler)
	e.DELETE("/api/currency/:id", deleteCurrencyHandler)
}
