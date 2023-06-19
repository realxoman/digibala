package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatGallery(t *testing.T) {
	e := echo.New()
	galleryRoutes(e)

	gallery := &models.Gallery{
		ID:        1,
		Name:      "armin",
		TypeFile:  "pdf",
		Size:      "a",
		Direction: "home",
		Hidden:    true,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(gallery)

	// response and writer
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/gallery", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Gallery)
	err := json.NewDecoder(resp.Body).Decode(ret)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, gallery, ret)
}

// Test delete gallery
func TestDeleteGallery(t *testing.T) {
	e := echo.New()
	galleryRoutes(e)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/gallery/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test Update gallery
func TestUpdateGallery(t *testing.T) {
	e := echo.New()
	galleryRoutes(e)

	shipping := &models.Gallery{}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/gallery", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Gallery)
	err := json.NewDecoder(resp.Body).Decode(ret)
	fmt.Println(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, shipping, ret)
}

func TestListGallery(t *testing.T) {
	e := echo.New()
	galleryRoutes(e)
	galleries := new([]*models.Gallery)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/galleries", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	resp := w.Result()
	ret := new([]*models.Gallery)
	fmt.Println(ret)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, galleries, ret)
}

// Get ById gallery
func TestGetByIDGallery(t *testing.T) {
	e := echo.New()
	galleryRoutes(e)
	gallery := &models.Gallery{
		ID: 1,
	}
	buf := new(bytes.Buffer)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/gallery/1", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Gallery)
	err := json.NewDecoder(resp.Body).Decode(ret)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, gallery, ret)
}
