package models

import "time"

type Promotion struct {
	ID          int
	Category    Category
	Name        string
	Description string
	Discount    float64
	ExpireDate  time.Time
}

type Category struct {
	ID   int
	Name string
}
