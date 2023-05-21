package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func tagRoutes(e *echo.Echo) {
	e.GET("/tags", listTagHandler)
	e.POST("/tags", createTagHandler)
	e.GET("/tags/:id", findTagHandler)
	e.DELETE("/tags/:id", deleteTagHandler)
	e.PUT("/tags", updateTagHandler)
}

func updateTagHandler(c echo.Context) error {
	tag := &models.Tag{}
	if err := c.Bind(tag); err != nil {
		return err
	}
	fmt.Println("Updating tags id:", tag.ID)
	return c.JSON(http.StatusOK, tag)
}

func deleteTagHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting tags id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createTagHandler(c echo.Context) error {
	tag := &models.Tag{}
	if err := c.Bind(tag); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tag)
}

func listTagHandler(c echo.Context) error {
	tags := []models.Tag{}
	return c.JSON(http.StatusOK, tags)
}

func findTagHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tag := &models.Tag{ID: id}
	return c.JSON(http.StatusOK, tag)
}
