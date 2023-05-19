package server

import (
	"digibala/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var currencies []models.Currency

func listCurrencyHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, currencies)
}

func getCurrencyByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	for _, currency := range currencies {
		if currency.ID == id {
			return c.JSON(http.StatusOK, currency)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Currency not found")
}

func createCurrencyHandler(c echo.Context) error {
	currency := new(models.Currency)
	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency data")
	}

	currency.ID = len(currencies) + 1
	currencies = append(currencies, *currency)

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

	for i := range currencies {
		if currencies[i].ID == id {
			currency.ID = id
			currencies[i] = *currency
			return c.JSON(http.StatusOK, currency)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Currency not found")
}

func deleteCurrencyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	for i := range currencies {
		if currencies[i].ID == id {
			currencies = append(currencies[:i], currencies[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Currency not found")
}

func currencyRoutes(e *echo.Echo) {
	e.GET("/api/currency", listCurrencyHandler)
	e.GET("/api/currency/:id", getCurrencyByIDHandler)
	e.POST("/api/currency", createCurrencyHandler)
	e.PUT("/api/currency/:id", updateCurrencyHandler)
	e.DELETE("/api/currency/:id", deleteCurrencyHandler)
}
