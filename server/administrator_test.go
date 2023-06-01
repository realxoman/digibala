package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAdministrator(t *testing.T) {
	e := echo.New()
	adminRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/admin/5", nil)
	r.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("Bad Request")
	}

	resp := w.Result()
	ret := new(models.StatusOK)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &models.StatusOK{OK: "OK"}, ret)
}

func TestCreateAdministrator(t *testing.T) {
	e := echo.New()
	adminRoutes(e)

	admin := &models.Administrator{UserId: 1, Id: 1, Rank: models.AdminRank{}, User: models.User{}}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(admin)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Bad Request")
	}

	resp := w.Result()
	ret := new(models.Administrator)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, admin, ret)
}

func TestUpdateAdministrator(t *testing.T) {
	e := echo.New()
	adminRoutes(e)

	admin := &models.Administrator{UserId: 1, Id: 1, Rank: models.AdminRank{}, User: models.User{}}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(admin)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/admin", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Bad Request")
	}

	resp := w.Result()
	ret := new(models.Administrator)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, admin, ret)
}

func TestCheckAdminHandler(t *testing.T) {
	e := echo.New()

	adminRoutes(e)

	userID := 1

	path := fmt.Sprintf("/admin/%d", userID)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)
	r.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Bad Request")
	}

	var admin models.Administrator
	err := json.NewDecoder(w.Body).Decode(&admin)
	if err != nil {
		t.Fatal(err)
	}

	expected := models.Administrator{UserId: userID}
	assert.Equal(t, expected, admin)
}

func TestListAdministrator(t *testing.T) {
	e := echo.New()
	adminRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/admin", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Bad Request")
	}

	resp := w.Result()
	var ret []models.Administrator
	err := json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []models.Administrator{}, ret)
}
