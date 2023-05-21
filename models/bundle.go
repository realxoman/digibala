package models

// language model comes here

type Bundle struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ProductIds  []int   `json:"product_ids"`
	Cost        float64 `json:"cost"`
	Discount    float64 `json:"discount"`
	IsActive    bool    `json:"is_active"`
}
