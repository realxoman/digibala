package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"digibala/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStockRoutes(t *testing.T) {
	e := echo.New()
	stockRoutes(e)

	t.Run("listStockHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/stock", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, listStockHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var stocks []models.Stock
			err := json.Unmarshal(rec.Body.Bytes(), &stocks)
			assert.NoError(t, err)
			assert.Equal(t, []models.Stock{}, stocks)
		}
	})

	t.Run("createStockHandler", func(t *testing.T) {
		reqBody := []byte(`{"product_id": 1, "quantity": 10}`)
		req := httptest.NewRequest(http.MethodPost, "/stock", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, createStockHandler(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			var stock models.Stock
			err := json.Unmarshal(rec.Body.Bytes(), &stock)
			assert.NoError(t, err)
			assert.Equal(t, 1, stock.ProductID)
			assert.Equal(t, 10, stock.Quantity)
		}
	})

	t.Run("findStockHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/stock/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/stock/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, findStockHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var id int
			err := json.Unmarshal(rec.Body.Bytes(), &id)
			assert.NoError(t, err)
			assert.Equal(t, 1, id)
		}
	})

	t.Run("updateStockHandler", func(t *testing.T) {
		reqBody := map[string]int{"quantity": 50}
		reqJSON, err := json.Marshal(reqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, "/stock/1", bytes.NewBuffer(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/stock/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, updateStockHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var quantity models.Stock
			err := json.Unmarshal(rec.Body.Bytes(), &quantity)
			assert.NoError(t, err)
			assert.Equal(t, 1, quantity.ProductID)
			assert.Equal(t, reqBody["quantity"], quantity.Quantity) // Get quantity value from JSON request body
		}
	})

	t.Run("deleteStockHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/stock/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/stock/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, deleteStockHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var status models.StatusOK
			err := json.Unmarshal(rec.Body.Bytes(), &status)
			assert.NoError(t, err)
			assert.Equal(t, "OK", status.OK)
		}
	})
}
