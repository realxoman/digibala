package service

import (
	models "admin/model"
	"admin/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = storage.GetConnection()
}

func Delete(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))

	admin := &models.Administrator{}
	result := db.First(admin, user_id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}

	result = db.Delete(admin)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}
	fmt.Println("Deleting Admin with User ID: ", user_id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func Update(c echo.Context) error {
	admin := &models.Administrator{}
	if err := c.Bind(admin); err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return err
	}
	admin.UserId = userID

	permLvl, err := strconv.Atoi(c.FormValue("permlvl"))
	if err != nil {
		return err
	}
	admin.PermLvl = permLvl

	existingAdmin := &models.Administrator{}
	result := db.First(existingAdmin, "user_id = ?", admin.UserId)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}

	existingAdmin.UserId = admin.UserId
	existingAdmin.PermLvl = admin.PermLvl

	result = db.Save(existingAdmin)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}
	fmt.Println("updating admin with user_id:", admin.UserId)

	return c.JSON(http.StatusOK, admin)
}

func Check(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	admin := &models.Administrator{UserId: user_id}
	result := db.First(admin, user_id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}
	fmt.Println("finding admin with user_id:", user_id)
	return c.JSON(http.StatusOK, admin)
}

func Add(c echo.Context) error {
	admin := &models.Administrator{}
	if err := c.Bind(admin); err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return err
	}
	admin.UserId = userID

	permLvl, err := strconv.Atoi(c.FormValue("permlvl"))
	if err != nil {
		return err
	}
	admin.PermLvl = permLvl

	result := db.Create(admin)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}
	fmt.Println("adding new valid admin")
	return c.JSON(http.StatusOK, admin)
}

func List(c echo.Context) error {
	admins := []models.Administrator{} // from db
	result := db.Find(&admins)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.StatusError{Error: fmt.Sprintf("%s", result.Error)})
	}
	fmt.Println("Listing all admins")
	return c.JSON(http.StatusOK, admins)
}
