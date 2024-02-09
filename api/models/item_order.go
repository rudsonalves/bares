package models

// ItemOrder structure of items ordered in the system
type ItemOrder struct {
	Id       int    `json:"id"`
	OrderId  int    `json:"orderId"`
	ItemId   int    `json:"itemId"`
	Amount   int    `json:"amount"`
	Comments string `json:"comments"`
}
