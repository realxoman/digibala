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

func TestHandlerListBrands(t *testing.T) {
	e := echo.New()
	brandRoutes(e)

	req := httptest.NewRequest(http.MethodGet, "/brands", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ListBrands(c)
	if err != nil {
		t.Fatalf("failed to handle ListBrands: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := []models.Brand{}
	actualResponse := []models.Brand{}
	err = json.Unmarshal(rec.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestHandlerCreateBrand(t *testing.T) {
	e := echo.New()
	brandRoutes(e)

	brand := &models.Brand{
		ID:          1,
		Name:        "Test Brand",
		Description: "Test Description",
		Logo:        "logo.png",
		CountryID:   123,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	expectedJSON, err := json.Marshal(brand)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(brand)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/brands", buf)
	r.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	responseBody := w.Body.String()
	assert.JSONEq(t, string(expectedJSON), responseBody)
}

func TestHandlerGetBrand(t *testing.T) {
	e := echo.New()
	brandRoutes(e)

	req := httptest.NewRequest(http.MethodGet, "/brands/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetBrand(c)
	if err != nil {
		t.Fatalf("failed to handle GetBrand: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody struct {
		ID int `json:"ID"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	expectedID := 1
	assert.Equal(t, expectedID, responseBody.ID)
}

func TestHandlerUpdateBrand(t *testing.T) {
	e := echo.New()
	brandRoutes(e)

	brand := &models.Brand{
		ID:   1,
		Name: "Updated Brand",
	}

	expectedJSON, err := json.Marshal(brand)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(brand)
	req := httptest.NewRequest(http.MethodPut, "/brands/1", buf)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = UpdateBrand(c)
	if err != nil {
		t.Fatalf("failed to handle UpdateBrand: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	responseBody := rec.Body.String()
	assert.JSONEq(t, string(expectedJSON), responseBody)
}

func TestHandlerDeleteBrand(t *testing.T) {
	e := echo.New()
	brandRoutes(e)

	req := httptest.NewRequest(http.MethodDelete, "/brands/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteBrand(c)
	if err != nil {
		t.Fatalf("failed to handle DeleteBrand: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := map[string]string{"OK": "OK"}
	actualResponse := make(map[string]string)
	err = json.Unmarshal(rec.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
