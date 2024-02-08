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
	Id       int       `json:"id"`
	UserId   int       `json:"userId"`
	DateHour time.Time `json:"dateHour"`
	Status   Status    `json:"status"`
}

func (p *Order) String() string {
	return fmt.Sprintf("id: %d; userId: %d; Date: %s; Status: %s",
		p.Id,
		p.UserId,
		p.DateHour.Format("02/01/06"),
		p.Status)
}
