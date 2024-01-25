package models

type Papel string

const (
	Cliente Papel = "cliente"
	Garcom  Papel = "garcom"
	Gerente Papel = "gerente"
)

type Usuario struct {
	UsuarioID int    `json:"usuarioID"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	SenhaHash string `json:"senhaHash"`
	Papel     Papel  `json:"papel"`
}
