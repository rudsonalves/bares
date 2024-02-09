package models

import "fmt"

type Role string

// Roles for users
const (
	Cliente Role = "cliente"
	Garcom  Role = "garcom"
	Gerente Role = "gerente"
	Admin   Role = "admin"
	Cozinha Role = "cozinha"
	Caixa   Role = "caixa"
)

// User struct for system users
type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Role         Role   `json:"role"`
}

func (u User) String() string {
	return fmt.Sprintf(
		"User {\n  Id:    %d,\n  Name:  %s,\n  Email: %s,\n  Senha: %s,\n  Role:  %v\n}",
		u.Id, u.Name, u.Email, u.PasswordHash, u.Role)
}
