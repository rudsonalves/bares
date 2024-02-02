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

// Pedido estrutura dos pedidos no sistema
type Pedido struct {
	PedidoID  int       `json:"pedidoID"`
	UsuarioID int       `json:"usuarioID"`
	DataHora  time.Time `json:"dataHora"`
	Status    Status    `json:"status"`
}

func (p *Pedido) String() string {
	return fmt.Sprintf("id: %d; userId: %d; Date: %s; Status: %s",
		p.PedidoID,
		p.UsuarioID,
		p.DataHora.Format("02/01/06"),
		p.Status)
}
