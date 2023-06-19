package models

import "gorm.io/gorm"

type TagType int

const (
	ACCOUNT int = iota
	PAYMENT
	RETURNPRODUCT
	SENDFEEDBACK
)

type FAQ struct {
	gorm.Model
	ID          int     `json:"id"`
	Question    string  `json:"question"`
	Answer      string  `json:"answer"`
	QuestionTag TagType `json:"questiontag"`
	//CreatedAt   time.Time `json:"created_at"`
	//UpdatedAt   time.Time `json:"updated_at"`
}
