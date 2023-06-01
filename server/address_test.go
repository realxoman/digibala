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

func TestCreateAddress(t *testing.T) {
	e := echo.New()
	addressRoutes(e)

	address := &models.Address{
		UserID:   5,
		Title:    "Home",
		Location: models.Location{1.1, 2.2},
		Address:  "Iran",
		No:       "12",
		Floor:    "7",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(address)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/address", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Address)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, address, ret)

}
