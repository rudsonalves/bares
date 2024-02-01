package handlers

import (
	"bares_api/models"
	"bares_api/services"

	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthHandler gerencia as requisições HTTP para Pedido.
type AuthHandler struct {
  Service *services.AuthService
}

// NewAuthHandler cria uma nova instância de AuthHandler.
func NewAuthHandler(service *services.AuthService) *AuthHandler {
  return &AuthHandler{
    Service: service,
  }
}

// LoginHandlers
func (handler *AuthHandler) LoginHandlers(w http.ResponseWriter, r *http.Request) {
  var credentials models.Credentials
  if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // Lógica de validação das credenciais
  papel, err := handler.Service.ValidarCredenciais(credentials)
  if err != nil {
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    return
  }

  // Gerar Token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "email": credentials.Email,
    "papel": papel,
    "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
  })

  tokenString, err := token.SignedString([]byte(secretKey))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // Retornar o token
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
