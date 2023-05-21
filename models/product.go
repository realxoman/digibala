package models

type Product struct {
	ID          int `json:"id"`
	Title       string
	Price       int
	Quantity    int
	Description string
	Image       string
}
