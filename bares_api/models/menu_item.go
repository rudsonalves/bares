package models

// MenuItem estrutura dos itens de menu do sistema
type MenuItem struct {
	ItemID    int     `json:"itemID"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	ImagemURL string  `json:"imagemURL"`
}
