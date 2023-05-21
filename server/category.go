package server

import (
	"digibala/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func categoryRoutes(e *echo.Echo) {
	e.GET("/category", listCategoryHandler)
	e.POST("/category", createCategoryHandler)
	e.GET("/category/:id", findCategoryHandler)
	e.DELETE("/category/:id", deleteCategoryHandler)
	e.PUT("/category", updateCategoryHandler)
}

func updateCategoryHandler(c echo.Context) error {
	category := &models.Category{}
	if err := c.Bind(category); err != nil {
		return err
	}
	fmt.Println("Updating category id:", category.Id)
	return c.JSON(http.StatusOK, category)
}

func deleteCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting category id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createCategoryHandler(c echo.Context) error {
	category := &models.Category{}
	if err := c.Bind(category); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, category)
}

func listCategoryHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []models.Category{})
}

func findCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category := &models.Category{Id: int32(id)}
	return c.JSON(http.StatusOK, category)
}
