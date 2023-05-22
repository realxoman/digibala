package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func paymentRoutes(e *echo.Echo) {
	e.GET("/payment", listPaymentHandler)
	e.POST("/payment", createPaymentHandler)
	e.GET("/payment/:id", findPaymentHandler)
	e.DELETE("/payment/:id", deletePaymentHandler)
	e.PUT("/payment", updatePaymentHandler)
}

func updatePaymentHandler(c echo.Context) error {
	payment := &models.PaymentMethod{}
	if err := c.Bind(payment); err != nil {
		return err
	}
	fmt.Println("Update Payment id:", payment.ID)
	return c.JSON(http.StatusOK, payment)

}

func deletePaymentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Remove payment :", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createPaymentHandler(c echo.Context) error {
	payment := &models.PaymentMethod{}
	if err := c.Bind(payment); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, payment)
}

func listPaymentHandler(c echo.Context) error {
	payments := []models.PaymentMethod{}
	return c.JSON(http.StatusOK, payments)
}

func findPaymentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	payment := models.PaymentMethod{ID: id}
	return c.JSON(http.StatusOK, payment)
}
