package server

import (
	"digibala/models"
	"net/http"
    //"time"
	"fmt"
	"strconv"

    "github.com/labstack/echo/v4"
)



// ListBrands retrieves a list of all brands
func ListBrands(c echo.Context) error {
	var brands []models.Brand
	return c.JSON(http.StatusOK, brands)
}

// create function 
func CreateBrand(c echo.Context) error {
	brand := new(models.Brand)
	if err := c.Bind(brand); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, brand)
}
//retrive function
func GetBrand(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	address := &models.Brand{ID: id}
	return c.JSON(http.StatusOK, address)
}

// UpdateBrand updates an existing brand
func UpdateBrand(c echo.Context) error {
	brand := &models.Brand{}
	if err := c.Bind(brand); err != nil {
		return err
	}
	fmt.Println("Updating brand id:", brand.ID)
	return c.JSON(http.StatusOK, brand)
}
// delete function 
func DeleteBrand(c echo.Context) error {
	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid brand ID")
	}
	fmt.Println("Deleting brand id:", brandID)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func brandRoutes(e *echo.Echo) {
	e.GET("/brands", ListBrands)
	e.POST("/brands", CreateBrand)
	e.GET("/brands/:id", GetBrand)
	e.DELETE("/brands/:id", DeleteBrand)
	e.PUT("/brands/:id", UpdateBrand)
}