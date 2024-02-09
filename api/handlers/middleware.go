package handlers

import (
	"bares_api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// secretKey used to sign tokens. Ideally, it should be complex and stored securely
// Fixed Keys with Secure Storage: Rather than generating the key randomly at startup,
// many systems choose to use a fixed key that is stored securely, such as in a
// password vault, a key management service, or injected into the runtime environment
// through protected environment variables. This maintains the security of the key
// while avoiding the problems associated with random generation at initialization.
const secretKey = "CPh@s?NU?<qHlb_T@dNK#tHE1r#6D1_iVBYgsRQ8h@mx3U!pfx7-uOE$I#l#!9w."

// AuthMiddleware checks the presence and validity of a JWT token.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			// Header format must be "Bearer <token>"
			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := headerParts[1]

			// Parse and validate the token
			token, err := jwt.Parse(tokenString,
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method")
					}
					return []byte(secretKey), nil
				},
			)

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Extrai o tipo de usuário (role) das claims do token
				roleStr, ok := claims["role"].(string)
				if !ok {
					http.Error(w, "Role claim must be a string", http.StatusBadRequest)
					return
				}
				userRole := models.Role(roleStr)
				fmt.Println("Claims: ", claims)

				// Implemente aqui a lógica de autorização baseada no tipo de usuário
				if !isAuthorized(userRole, r.URL.Path, r.Method) {
					http.Error(w, "Access denied", http.StatusForbidden)
					return
				}

				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
		},
	)
}

// isAuthorized verifica se o usuário tem permissão para acessar o recurso com base no seu papel
func isAuthorized(userRole models.Role, path, method string) bool {
	// Define padrões de caminho para identificar ações específicas
	// Nota: estas são strings simplificadas, você pode precisar ajustar ou usar regex para caminhos mais complexos
	isUserCreation := path == "/users" && method == "POST"
	isOrderCreation := path == "/orders" && method == "POST"
	isOrderUpdate := strings.HasPrefix(path, "/orders/") && method == "PUT"
	isItemOrderUpdate := strings.HasPrefix(path, "/itemOrder/") && method == "PUT"

	switch userRole {
	case models.Garcom:
		// Garçom pode criar clientes e pedidos, e atualizar pedidos e itemOrder
		if isUserCreation || isOrderCreation || isOrderUpdate || isItemOrderUpdate {
			return true
		}
	case models.Cliente:
		// Cliente pode criar pedidos
		if isOrderCreation {
			return true
		}
	case models.Cozinha:
		// Cozinha pode atualizar pedidos
		if isOrderUpdate {
			return true
		}
	case models.Admin, models.Gerente:
		// Admin e Gerente presumivelmente têm acesso total para simplificação, ajuste conforme necessário
		return true
	}

	// Por padrão, nega acesso se nenhuma das condições acima for atendida
	return false
}
