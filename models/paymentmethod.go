package models

type DargahType struct {
	Saman    string
	AP       string
	Pasargad string
}

type PaymentMethod struct {
	ID     int        `json:"id"`
	UserID int        `json:"user_id"`
	Status string     `json:"status"`
	Dargah DargahType `json:"dargah"`
}
