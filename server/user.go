package server

import (
	"digibala/models"

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
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, newUser)
}
func editUserHandler(c echo.Context) error {
	if err := c.Bind(models.User{}); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, models.User{})
}
func deleteUserHandler(c echo.Context) error {
	id := c.Param("id")
	_ = id
	return c.JSON(http.StatusOK, map[string]interface{}{
		
		"user": models.User{},
	})
}

func init() {

	e.GET("/users", getAllUsersHandler)
	e.GET("/users/:id", getUserHandler)
	e.POST("/users", createUserHandler)
	e.PUT("/users/:id", editUserHandler)
	e.DELETE("/users/:id", deleteUserHandler)

}
