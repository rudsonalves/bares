package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// secretKey usado para assinar tokens. Idealmente, deve ser complexo e armazenado com segurança
const secretKey = "your_secret_key"

// AuthMiddleware verifica a presença e validade de um token JWT.
func AuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Extrair o token do cabeçalho Authorization
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
      http.Error(w, "Authorization header is required", http.StatusUnauthorized)
      return
    }

    // O formato do cabeçalho deve ser "Bearer <token>"
    headerParts := strings.Split(authHeader, " ")
    if len(headerParts) != 2 || headerParts[0] != "Bearer" {
      http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
      return
    }

    tokenString := headerParts[1]

    // Parse e valida o token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method")
      }
      return []byte(secretKey), nil
    })

    if err != nil {
      http.Error(w, "Invalid token", http.StatusUnauthorized)
      return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      // Adicionar informações do usuário ao contexto da requisição, se necessário
      // Exemplo: ctx := context.WithValue(r.Context(), "userID", claims["userID"])
      // r = r.WithContext(ctx)
      fmt.Println("Claims: ", claims)

      next.ServeHTTP(w, r)
    } else {
      http.Error(w, "Invalid token", http.StatusUnauthorized)
    }
  })
}
