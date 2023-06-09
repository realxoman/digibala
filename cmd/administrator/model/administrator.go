package models

import "gorm.io/gorm"

type Administrator struct {
	gorm.Model
	UserId  int `json:"user_id"`
	PermLvl int `json:"permlvl"`
}

type StatusOK struct {
	OK string
}

type StatusError struct {
	Error string
}
