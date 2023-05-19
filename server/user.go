package server

import (
	"digibala/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// user routes comes here

func getAllUsersHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get All Users")
}
func getUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get Single User")
}

func createUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "create User") 
}
func editUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Edit user successfully")
}
func deleteUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Delete User successfully")
}

func init() {
	fmt.Println(models.User{})
	e.GET("/users", getAllUsersHandler)
	e.GET("/users/:id", getUserHandler)
	e.POST("/users", createUserHandler)
	e.PUT("/users/:id", editUserHandler)
	e.DELETE("/users/:id", deleteUserHandler)
}
