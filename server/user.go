package server

import (
	"digibala/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// user routes comes here

func getAllUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.User{})
}
func getUserHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.User{})
}

func createUserHandler(c echo.Context) error {
	newUser := new(models.User)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, newUser)
}
func editUserHandler(c echo.Context) error {
	if err := c.Bind(models.User{}); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, models.User{})
}
func deleteUserHandler(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"id":      id,
		"user":    models.User{},
	})
}

func init() {
	fmt.Println(models.User{})
	e.GET("/users", getAllUsersHandler)
	e.GET("/users/:id", getUserHandler)
	e.POST("/users", createUserHandler)
	e.PUT("/users/:id", editUserHandler)
	e.DELETE("/users/:id", deleteUserHandler)
}
