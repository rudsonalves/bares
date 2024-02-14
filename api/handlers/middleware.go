package handlers

import (
	"bares_api/models"
	"fmt"
	"log"
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
	log.Println("AuthMiddleware: starting...")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Printf("AuthMiddleware 0: Authorization header is required")
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			// Header format must be "Bearer <token>"
			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				log.Printf("AuthMiddleware 1: Invalid Authorization header format")
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := headerParts[1]

			// Parse and validate the token
			token, err := jwt.Parse(tokenString,
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						log.Printf("AuthMiddleware 2: unexpected signing method")
						return nil, fmt.Errorf("unexpected signing method")
					}
					return []byte(secretKey), nil
				},
			)

			if err != nil {
				log.Printf("AuthMiddleware 3: Invalid token")
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Extracts the user type (role) from the token claims
				roleStr, ok := claims["role"].(string)
				if !ok {
					log.Printf("AuthMiddleware 4: Role claim must be a string")
					http.Error(w, "Role claim must be a string", http.StatusBadRequest)
					return
				}
				userRole := models.Role(roleStr)
				fmt.Println("Claims: ", claims)

				// Implement authorization logic based on user type here
				if !isAuthorized(userRole, r.URL.Path, r.Method) {
					log.Printf("AuthMiddleware 4: Access denied")
					http.Error(w, "Access denied", http.StatusForbidden)
					return
				}

				next.ServeHTTP(w, r)
			} else {
				log.Printf("AuthMiddleware 5: Invalid token")
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
		},
	)
}

// isAuthorized checks if the user is allowed to access the resource based on their role
func isAuthorized(userRole models.Role, path, method string) bool {
	// Defines path patterns to identify specific actions
	// Note: these are simplified strings, you may need to adjust or use regex for more complex paths
	isUserCreation := path == "/users" && method == "POST"
	isUserUpdate := strings.Contains(path, "/users") && method == "PUT"
	// isUserDelete := strings.Contains(path, "/users") && method == "DELETE"
	isOrderCreation := path == "/orders" && method == "POST"
	isOrderUpdate := strings.HasPrefix(path, "/orders/") && method == "PUT"
	isItemOrderUpdate := strings.HasPrefix(path, "/itemOrder/") && method == "PUT"

	switch userRole {
	case models.Garcom:
		// Waiter can create customers and orders, and update orders and itemOrder
		if isUserCreation || isOrderCreation || isOrderUpdate || isItemOrderUpdate || isUserUpdate {
			return true
		}
	case models.Cliente:
		// Customer can create orders
		if isOrderCreation {
			return true
		}
	case models.Cozinha:
		// Kitchen can update orders
		if isOrderUpdate {
			return true
		}
	case models.Gerente:
		// Admin and Manager presumably have full access for simplification, adjust as needed
		return true
	case models.Admin:
		// Admin and Manager presumably have full access for simplification, adjust as needed
		return true
	}

	// By default deny access if none of the above conditions are met
	return false
}
