package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var voucher = &models.Voucher{
	ID:          1,
	IsActive:    true,
	Code:        "YaldaBala",
	Type:        "Percentage",
	Product:     models.ProductVoucher{IsActive: true, ProductsId: []int{1, 2, 3}},
	User:        models.UserVoucher{IsActive: true, UsersId: []int{1, 2, 3}},
	ExpiredTime: time.Date(2023, 11, 17, 20, 34, 58, 651387237, time.UTC),
}

func TestListVoucher(t *testing.T) {
	e := echo.New()
	voucherRoutes(e)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/voucher", nil)
	request.Header.Set("content-type", "application/json")
	e.ServeHTTP(recorder, request)

	response := recorder.Result()
	ret := []*models.Voucher{}
	err := json.NewDecoder(response.Body).Decode(&ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []*models.Voucher{}, ret)
}

func TestGetVoucher(t *testing.T) {
	e := echo.New()
	voucherRoutes(e)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(voucher)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/voucher/1", buf)
	request.Header.Set("content-type", "application/json")
	e.ServeHTTP(recorder, request)

	resp := recorder.Result()
	ret := new(models.Voucher)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, &models.Voucher{ID: 1}, ret)
}

func TestCreateVoucher(t *testing.T) {
	e := echo.New()
	voucherRoutes(e)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(voucher)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/voucher", buf)
	request.Header.Set("content-type", "application/json")
	e.ServeHTTP(recorder, request)

	response := recorder.Result()
	ret := new(models.Voucher)
	err := json.NewDecoder(response.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)
	assert.Equal(t, voucher, ret)

}

func TestDeleteVoucher(t *testing.T) {
	e := echo.New()
	voucherRoutes(e)

	status := &models.StatusOK{
		OK: "OK",
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/api/v1/voucher/5", nil)
	request.Header.Set("content-type", "application/json")
	e.ServeHTTP(recorder, request)

	resp := recorder.Result()
	ret := new(models.StatusOK)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, status, ret)
}

func TestUpdateVoucher(t *testing.T) {
	e := echo.New()
	voucherRoutes(e)

	voucher := &models.Voucher{
		ID: 0,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/api/v1/voucher/5", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Voucher)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, voucher, ret)
}
