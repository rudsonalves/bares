package models

// Credentials represent the user's login credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
