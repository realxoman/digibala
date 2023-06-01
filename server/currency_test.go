package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCurrencyHandler(t *testing.T) {
	e := echo.New()
	currencyRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/currency", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new([]*models.Currency)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, new([]*models.Currency), ret)
}

func TestGetCurrencyByIDHandler(t *testing.T) {
	e := echo.New()
	currencyRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/currency/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Currency)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, &models.Currency{ID: 1}, ret)
}

func TestCreateCurrencyHandler(t *testing.T) {
	e := echo.New()
	currencyRoutes(e)

	currency := &models.Currency{
		ID:     5,
		Code:   "IR",
		Name:   "Rial",
		Symbol: "IR",
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(currency)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/currency", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Currency)
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
	assert.Equal(t, currency, ret)
}
func TestUpdateCurrencyHandler(t *testing.T) {
	e := echo.New()
	currencyRoutes(e)

	currency := &models.Currency{
		ID:     1,
		Code:   "IR",
		Name:   "Rial",
		Symbol: "IR",
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(currency)
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/api/currency/1", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Currency)
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, currency, ret)
}
func TestDeleteCurrencyHandler(t *testing.T) {
	e := echo.New()
	currencyRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/api/currency/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
