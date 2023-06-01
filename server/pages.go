package server

import (
	"digibala/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func pagesRoute(e *echo.Echo) {
	e.GET("/about", aboutHandler)
}

func aboutHandler(c echo.Context) error {
	teacher := models.Teacher{
		FirstName: "Mostafa",
		LastName:  "Solati",
		Email:     "mostafa.solati@gmail.com",
	}
	data := map[string]interface{}{
		"teacher":  teacher,
		"subjects": []string{"context", "design pattern", "decorator", "adapter", "proxy"},
		"finished": false,
	}
	return c.Render(http.StatusOK, "about.html", data)
}
