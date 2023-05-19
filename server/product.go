package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var products []models.Product

func productRoutes(e *echo.Echo) {
	fmt.Println(models.User{})
	e.GET("/product", indexProductHandler)
	e.GET("/product/:id", FindProductHandler)
	e.POST("/product", createProductHandler)
	e.PUT("/product/:id", updateProductHandler)
	e.DELETE("/product/:id", deleteProductHandler)
}

// user routes comes here
func indexProductHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}
func FindProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, product := range products {
		if product.ID == id {
			return c.JSON(http.StatusOK, product)
		}
	}
	return c.JSON(http.StatusNotFound, "product not found")
}
func createProductHandler(c echo.Context) error {
	product := &models.Product{}

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	products = append(products, *product)
	return c.JSON(http.StatusCreated, product)
}
func updateProductHandler(c echo.Context) error {
	product := &models.Product{}
	id, _ := strconv.Atoi(c.Param("id"))
	var index int
	find := false
	for i, product := range products {
		fmt.Println(product.ID, product.ID == id)
		if product.ID == id {
			index = i
			find = true
		}
	}
	if !find {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	products[index] = *product

	return c.JSON(http.StatusOK, products[index])
}
func deleteProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var index int
	find := false
	for i, product := range products {
		fmt.Println(product.ID, product.ID == id)
		if product.ID == id {
			index = i
			find = true
		}
	}
	if !find {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	products = append(products[:index], products[index+1:]...)

	return c.JSON(http.StatusOK, "product deleted successfully")
}
