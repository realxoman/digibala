package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func adminRoutes(e *echo.Echo) {
	e.GET("/admin", listAdminHandler)
	e.POST("/admin", addNewAdminHandler)
	e.GET("/admin/:user_id", checkAdminHandler)
	e.DELETE("/admin/:user_id", deleteAdminHandler)
	e.PUT("/admin", updateAdminHandler)
}

func deleteAdminHandler(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	fmt.Println("Deleting Admin with User ID: ", user_id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func updateAdminHandler(c echo.Context) error {
	admin := &models.Administrator{}
	if err := c.Bind(admin); err != nil {
		return err
	}
	fmt.Println("updating admin with user_id:", admin.User_id)
	return c.JSON(http.StatusOK, admin)
}

func checkAdminHandler(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	admin := &models.Administrator{User_id: user_id}
	fmt.Println("finding admin with user_id:", user_id)
	return c.JSON(http.StatusOK, admin)
}

func addNewAdminHandler(c echo.Context) error {
	admin := &models.Administrator{}
	if err := c.Bind(admin); err != nil {
		return err
	}
	fmt.Println("adding new valid admin")
	return c.JSON(http.StatusOK, admin)
}

func listAdminHandler(c echo.Context) error {
	admins := []models.Administrator{} // from db
	fmt.Println("Listing all admins")
	return c.JSON(http.StatusOK, admins)
}
