package server

import (
	"encoding/json"
	"faq/data"
	"faq/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

func init() {
	db = data.OpenDatabase()
}

func ListAllFAQService(c echo.Context) error {
	//query := `select id, question, answer, questiontag , created_at, updated_at from faqs`
	//
	//rows, err := db.Query(query)
	//defer rows.Close()
	//if err != nil {
	//	return nil, err
	//}
	//
	//var faqs []*FAQ
	//
	//for rows.Next() {
	//	var faq FAQ
	//	err := rows.Scan(
	//		&faq.ID,
	//		&faq.Question,
	//		&faq.Answer,
	//		&faq.QuestionTag,
	//		&faq.CreatedAt,
	//		&faq.UpdatedAt,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	faqs = append(faqs, &faq)
	//}
	faqs := []models.FAQ{}
	db.Find(&faqs)
	return c.JSON(http.StatusOK, faqs)
}
func FindSpecificFAQService(c echo.Context) error {
	//query := `select id, question, answer, questiontag , created_at, updated_at from faqs where id = $1`
	//row := db.QueryRow(query, id)
	//
	//var faq FAQ
	//err := row.Scan(
	//	&faq.ID,
	//	&faq.Question,
	//	&faq.Answer,
	//	&faq.QuestionTag,
	//	&faq.CreatedAt,
	//	&faq.UpdatedAt,
	//)
	id, _ := strconv.Atoi(c.Param("id"))
	faq := &models.FAQ{}
	faq.ID = id
	db.First(faq, id)
	return c.JSON(http.StatusOK, faq)
}
func CreateFAQService(c echo.Context) error {
	//var generatedID int
	//
	//query := `insert into faqs (question, answer, questiontag, created_at, updated_at)
	//	values ($1, $2, $3, $4, $5) returning id`
	//
	//err := db.QueryRow(query,
	//	faq.Question,
	//	faq.Answer,
	//	faq.QuestionTag,
	//	time.Now(),
	//	time.Now(),
	//).Scan(&generatedID)
	faq := &models.FAQ{}
	if err := c.Bind(faq); err != nil {
		return err
	}
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return err
	}
	faq.ID, _ = strconv.Atoi(jsonBody["id"].(string))
	tmp, _ := strconv.Atoi(jsonBody["questiontag"].(string))
	faq.QuestionTag = models.TagType(tmp)
	faq.Question = jsonBody["question"].(string)
	faq.Answer = jsonBody["answer"].(string)
	db.Create(faq)
	return c.JSON(http.StatusOK, faq)
}
func DeleteFAQService(c echo.Context) error {
	//query := `delete from faqs where id = $1`
	//_, err := db.Exec(query, id)
	id, _ := strconv.Atoi(c.Param("id"))
	faq := &models.FAQ{}
	db.First(faq, id)
	fmt.Printf("FAQ with id %d deleted", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func UpdateSpecificFAQService(c echo.Context) error {
	//query := `update faqs set
	//	question = $1,
	//	answer = $2,
	//	questiontag = $3,
	//	updated_at = $5
	//	where id = $6
	//`
	//
	//_, err := db.Exec(query,
	//	f.Question,
	//	f.Answer,
	//	f.QuestionTag,
	//	time.Now(),
	//	f.ID,
	//)
	id, _ := strconv.Atoi(c.Param("id"))
	faq := &models.FAQ{}
	faq.ID = id
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return err
	}
	faq.ID, _ = strconv.Atoi(jsonBody["id"].(string))
	tmp, _ := strconv.Atoi(jsonBody["questiontag"].(string))
	faq.QuestionTag = models.TagType(tmp)
	faq.Question = jsonBody["question"].(string)
	faq.Answer = jsonBody["answer"].(string)
	savedFAQ := &models.FAQ{}
	db.First(savedFAQ, "id = ?", id)
	savedFAQ.Question = faq.Question
	savedFAQ.Answer = faq.Answer
	savedFAQ.QuestionTag = faq.QuestionTag
	db.Save(savedFAQ)
	fmt.Printf("FAQ with id %d updated", id)
	return c.JSON(http.StatusOK, faq)
}
