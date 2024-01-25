package models

import "time"

type Status string

const (
	Recebido   Status = "recebido"
	Preparando Status = "preparando"
	Pronto     Status = "pronto"
	Entregue   Status = "entregue"
)

type Pedido struct {
	PedidoID  int       `json:"pedidoID"`
	UsuarioID int       `json:"usuarioID"`
	DataHora  time.Time `json:"dataHora"`
	Status    Status    `json:"status"`
}
