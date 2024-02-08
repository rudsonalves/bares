package models

// Credentials representa as credenciais do usuário para login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
