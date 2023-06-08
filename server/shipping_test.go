package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateShipping(t *testing.T) {
	e := echo.New()
	shippingRoutes(e)

	shipping := &models.Shipping{
		ProductID: 1,
		AddressID: 1,
		Timestamp: "2023-05-31 10:57:00",
		Type:      "Urgent",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(shipping)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/shipping", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Shipping)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, shipping, ret)
}

func TestListShipping(t *testing.T) {
	e := echo.New()
	shippingRoutes(e)

	shipping := new([]*models.Shipping)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/shipping", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new([]*models.Shipping)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, shipping, ret)
}

func TestFindShipping(t *testing.T) {
	e := echo.New()
	shippingRoutes(e)

	shipping := &models.Shipping{
		ID: 1,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/shipping/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Shipping)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, shipping, ret)
}

func TestDeleteShipping(t *testing.T) {
	e := echo.New()
	shippingRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/shipping/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestUpdateShipping(t *testing.T) {
	e := echo.New()
	shippingRoutes(e)

	shipping := &models.Shipping{}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/shipping", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Shipping)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, shipping, ret)
}
