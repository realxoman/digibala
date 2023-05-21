package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func bundleRoutes(e *echo.Echo) {
	e.GET("/bundle", listBundleHandler)
	e.GET("/bundle/:id", findBundleHandler)
	e.POST("/bundle", createBundleHandler)
	e.DELETE("/bundle/:id", deleteBundleHandler)
	e.PUT("/bundle", updateBundleHandler)
}

func updateBundleHandler(c echo.Context) error {
	bundle := &models.Bundle{}
	if err := c.Bind(bundle); err != nil {
		return err
	}
	fmt.Println("Updating bundle id:", bundle.ID)
	return c.JSON(http.StatusOK, bundle)
}

func deleteBundleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting bundle id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createBundleHandler(c echo.Context) error {
	bundle := &models.Bundle{}
	if err := c.Bind(bundle); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, bundle)
}

func listBundleHandler(c echo.Context) error {
	bundles := []models.Bundle{}
	return c.JSON(http.StatusOK, bundles)
}

func findBundleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bundle := &models.Bundle{ID: id}
	return c.JSON(http.StatusOK, bundle)
}
