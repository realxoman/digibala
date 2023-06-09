package server

import (
	"encoding/json"
	"fmt"
	"promotion/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type PromotionService struct {
	db *gorm.DB
}

func NewPromotionService(db *gorm.DB) *PromotionService {
	return &PromotionService{db: db}
}

func (s *PromotionService) listPromotionHandler(c echo.Context) error {
	var promotions []models.Promotion
	if err := s.db.Find(&promotions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, promotions)
}

func (s *PromotionService) findPromotionHandler(c echo.Context) error {
	promotionID := c.Param("id")
	id, _ := strconv.Atoi(promotionID)

	var promotion models.Promotion
	if err := s.db.First(&promotion, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "promotion not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, promotion)
}

func (s *PromotionService) deletePromotionHandler(c echo.Context) error {
	promotionID := c.Param("id")
	id, _ := strconv.Atoi(promotionID)

	var promotion models.Promotion
	if err := s.db.First(&promotion, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Promotion not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := s.db.Delete(&promotion).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *PromotionService) createPromotionHandler(c echo.Context) error {
	var promotion models.Promotion

	if err := c.Bind(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid promotion data")
	}

	if err := s.db.Create(&promotion).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, promotion)
}

func (s *PromotionService) updatePromotionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var promotion models.Promotion
	if err := c.Bind(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid promotion data")
	}

	if err := s.db.Model(&models.Promotionn{}).Where("id = ?", id).Updates(promotion).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "promotion not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	promotion.ID = uint(id)
	return c.JSON(http.StatusOK, promotion)
}

func RegisterRoutes(e *echo.Echo, promotionService *PromotionService) {
	e.GET("/promotion", promotionService.listPromotionHandler)
	e.GET("/promotion/:id", promotionService.findPromotionHandler)
	e.DELETE("/promotion/:id", promotionService.deletePromotionHandler)
	e.PUT("/promotion/category/:name", promotionService.createPromotionHandler)
	e.PATCH("/promotion/:id", promotionService.updatePromotionHandler)
}
