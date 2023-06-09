package models

type CurrencyRequest struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Exchange float64 `json:"exchange"`
}
