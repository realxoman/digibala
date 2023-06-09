package server

import (
	server "faq/server"
	"github.com/labstack/echo/v4"
)

func FaqRoutes(e *echo.Echo) {
	e.GET("/faqs", listAllFAQHandler)
	e.GET("/faq/:id", findSpecificFAQHandler)
	e.POST("/faq", createFAQHandler)
	e.DELETE("/faqs/:id", deleteFAQHandler)
	e.PUT("/faq/:id", updateSpecificFAQHandler)
}
func listAllFAQHandler(c echo.Context) error {
	return server.ListAllFAQService(c)
}
func findSpecificFAQHandler(c echo.Context) error {
	return server.FindSpecificFAQService(c)
}
func createFAQHandler(c echo.Context) error {
	return server.CreateFAQService(c)
}
func deleteFAQHandler(c echo.Context) error {
	return server.DeleteFAQService(c)
}
func updateSpecificFAQHandler(c echo.Context) error {
	return server.UpdateSpecificFAQService(c)
}
