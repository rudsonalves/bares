package models

// ItemOrder estrutura dos itens pedidos no sistema
type ItemOrder struct {
	Id       int    `json:"id"`
	OrderId  int    `json:"orderId"`
	ItemId   int    `json:"itemId"`
	Amount   int    `json:"amount"`
	Comments string `json:"comments"`
}
