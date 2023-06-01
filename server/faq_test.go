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

func TestCreateFAQ(t *testing.T) {
	e := echo.New()
	faqRoutes(e)

	faq := &models.FAQ{
		ID:          1,
		Question:    "Question 1",
		Answer:      "Answer 1",
		QuestionTag: []models.TagType{1, 2},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(faq)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/faq", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.FAQ)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, faq, ret)
}
func TestListAllFAQ(t *testing.T) {
	e := echo.New()
	faqRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/faqs", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new([]*models.FAQ)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, new([]*models.FAQ), ret)
}
func TestFindSpecificFAQ(t *testing.T) {
	e := echo.New()
	faqRoutes(e)

	SpecFaq := &models.FAQ{
		ID:          1,
		Question:    "Question 1",
		Answer:      "Answer 1",
		QuestionTag: []models.TagType{},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(SpecFaq)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/faq/1", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.FAQ)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, SpecFaq, ret)
}
func TestDeleteFAQ(t *testing.T) {
	e := echo.New()
	faqRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/faqs/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
func TestUpdateFAQ(t *testing.T) {
	e := echo.New()
	faqRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/faq/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
