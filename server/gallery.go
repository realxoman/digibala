package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func galleryRoutes(e *echo.Echo) {
	e.GET("/galleries", listGalleryHandler)
	e.POST("/gallery", createGalleryHandler)
	e.GET("/gallery/:id", findGalleryHandler)
	e.DELETE("/gallery/:id", deleteGalleryHandler)
	e.PUT("/gallery", updateGalleryHandler)
}

func updateGalleryHandler(c echo.Context) error {
	gallery := &models.Gallery{}
	//TODO logic Service
	if err := c.Bind(gallery); err != nil {
		return err
	}
	fmt.Println("Updating address id:", gallery.ID)
	return c.JSON(http.StatusOK, gallery)
}

func deleteGalleryHandler(c echo.Context) error {
	//TODO logic Service
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting address id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func createGalleryHandler(c echo.Context) error {
	gallery := &models.Gallery{}
	if err := c.Bind(gallery); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, gallery)
}

func listGalleryHandler(c echo.Context) error {
	var galleries []*models.Gallery
	return c.JSON(http.StatusOK, galleries)
}

func findGalleryHandler(c echo.Context) error {
	//TODO logic Service
	id, _ := strconv.Atoi(c.Param("id"))
	gallery := &models.Gallery{ID: int64(id)}
	return c.JSON(http.StatusOK, gallery)
}
