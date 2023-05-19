package server

import (
	"digibala/models"
	"net/http"
    "time"
	//"fmt"
	"strconv"

    "github.com/labstack/echo/v4"
)

var brands []models.Brand
var brandIDCounter int


// ListBrands retrieves a list of all brands
func ListBrands(c echo.Context) error {
	return c.JSON(http.StatusOK, brands)
}

// create function 
func CreateBrand(c echo.Context) error {
	brand := new(models.Brand)
	if err := c.Bind(brand); err != nil {
		return err
	}
	brand.ID = brandIDCounter
	brand.CreatedAt = time.Now()
	brand.UpdatedAt = time.Now()
	brandIDCounter++

	brands = append(brands, *brand)

	return c.JSON(http.StatusCreated, brand)
}
//retrive function
func GetBrand(c echo.Context) error {
	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid brand ID")
	}

	for _, brand := range brands {
		if brand.ID == brandID {
			return c.JSON(http.StatusOK, brand)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Brand not found")
}

// UpdateBrand updates an existing brand
func UpdateBrand(c echo.Context) error {
	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid brand ID")
	}

	var updatedBrand models.Brand
	found := false

	for i, brand := range brands {
		if brand.ID == brandID {
			if err := c.Bind(&updatedBrand); err != nil {
				return err
			}
			updatedBrand.ID = brand.ID
			updatedBrand.CreatedAt = brand.CreatedAt
			updatedBrand.UpdatedAt = time.Now()
			brands[i] = updatedBrand
			found = true
			break
		}
	}

	if found {
		return c.JSON(http.StatusOK, updatedBrand)
	}

	return echo.NewHTTPError(http.StatusNotFound, "Brand not found")
}
// delete function 
func DeleteBrand(c echo.Context) error {
	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid brand ID")
	}

	for i, brand := range brands {
		if brand.ID == brandID {
			brands = append(brands[:i], brands[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Brand not found")
}

func brandRoutes(e *echo.Echo) {
	e.GET("/brands", ListBrands)
	e.POST("/brands", CreateBrand)
	e.GET("/brands/:id", GetBrand)
	e.DELETE("/brands/:id", DeleteBrand)
	e.PUT("/brands/:id", UpdateBrand)
}