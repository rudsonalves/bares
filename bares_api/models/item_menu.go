package models

type ItemMenu struct {
	ItemID    int     `json:"itemID"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	ImagemURL string  `json:"imagemURL"`
}
