package handler

import (
	"admin/service"
	"admin/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = storage.GetConnection()
}

func AdminRoutes(e *echo.Echo) {
	e.GET("/api/admin", listAdminHandler)
	e.POST("/api/admin", addNewAdminHandler)
	e.GET("/api/admin/:user_id", checkAdminHandler)
	e.DELETE("/api/admin/:user_id", deleteAdminHandler)
	e.PUT("/api/admin", updateAdminHandler)
}

func deleteAdminHandler(c echo.Context) error {
	return service.Delete(c)
}

func updateAdminHandler(c echo.Context) error {
	return service.Update(c)
}

func checkAdminHandler(c echo.Context) error {
	return service.Check(c)
}

func addNewAdminHandler(c echo.Context) error {
	return service.Add(c)
}

func listAdminHandler(c echo.Context) error {
	return service.List(c)
}
