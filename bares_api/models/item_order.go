package models

// ItemOrder estrutura dos itens pedidos no sistema
type ItemOrder struct {
	ItemPedidoID int    `json:"itemPedidoID"`
	PedidoID     int    `json:"pedidoID"`
	ItemID       int    `json:"itemID"`
	Quantidade   int    `json:"quantidade"`
	Observacoes  string `json:"observacoes"`
}
