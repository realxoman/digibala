package models

import (
	"gorm.io/gorm"
	"time"
)

type Promotion struct {
	gorm.Model
	ID          int
	Category    Category
	Name        string
	Description string
	Discount    float64
	ExpireDate  time.Time
}

type Category struct {
	gorm.Model
	ID   int
	Name string
}
