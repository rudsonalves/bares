package models

import (
	"fmt"
	"time"
)

type Status string

const (
	Recebido   Status = "recebido"
	Preparando Status = "preparando"
	Pronto     Status = "pronto"
	Entregue   Status = "entregue"
)

// Order estrutura dos pedidos no sistema
type Order struct {
	PedidoID  int       `json:"pedidoID"`
	UsuarioID int       `json:"usuarioID"`
	DataHora  time.Time `json:"dataHora"`
	Status    Status    `json:"status"`
}

func (p *Order) String() string {
	return fmt.Sprintf("id: %d; userId: %d; Date: %s; Status: %s",
		p.PedidoID,
		p.UsuarioID,
		p.DataHora.Format("02/01/06"),
		p.Status)
}
