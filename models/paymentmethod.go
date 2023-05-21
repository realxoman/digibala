package models

type DargahType int

const (
	Saman DargahType = iota
	AP
	Pasargad
)

type PaymentMethod struct {
	ID     int        `json:"id"`
	UserID int        `json:"user_id"`
	Status string     `json:"status"`
	Dargah DargahType `json:"dargah"`
}
