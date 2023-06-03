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

func TestCreatePayment(t *testing.T) {
	e := echo.New()
	paymentRoutes(e)

	payment := &models.PaymentMethod{
		UserID: 5,
		Status: "Failed",
		Dargah: 1,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(payment)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/payment", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.PaymentMethod)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, payment, ret)
}

func TestListPayment(t *testing.T) {
	e := echo.New()
	paymentRoutes(e)

	payments := []models.PaymentMethod{}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/payment", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	var ret []models.PaymentMethod
	err := json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, payments, ret)
}

func TestFindPayment(t *testing.T) {
	e := echo.New()
	paymentRoutes(e)

	payments := &models.PaymentMethod{ID: 3}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/payment/3", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.PaymentMethod)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, payments, ret)
}

func TestDeletePayment(t *testing.T) {
	e := echo.New()
	paymentRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/payment/3", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdatePayment(t *testing.T) {
	e := echo.New()
	paymentRoutes(e)

	payments := &models.PaymentMethod{}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/payment", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.PaymentMethod)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, payments, ret)
}
