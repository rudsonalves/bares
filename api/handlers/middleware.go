package handlers

import (
	"bares_api/models"
	"fmt"
	"log"
	"net/http"
	"regexp"
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

type Permission struct {
	Methods map[string]bool
	Roles   map[models.Role]bool
}

var staticsPermissions = map[string]map[string][]models.Role{
	"/api/users": {
		"POST": {models.Garcom, models.Gerente, models.Admin},
		"GET":  {models.Gerente, models.Admin, models.Caixa},
	},
	"/api/menuitem": {
		"POST": {models.Gerente, models.Admin},
		"GET":  {models.Gerente, models.Admin, models.Cozinha},
	},
	"/api/order": {
		"POST": {models.Cliente, models.Garcom, models.Gerente, models.Admin},
		"GET":  {models.Caixa, models.Gerente, models.Gerente, models.Admin, models.Cozinha, models.Caixa},
	},
	"/api/itemorder": {
		"POST": {models.Cliente, models.Garcom, models.Gerente, models.Admin},
	},
}

func CheckPermission(path, method string, userRole models.Role) bool {
	if methodRoles, exists := staticsPermissions[path][method]; exists {
		for _, role := range methodRoles {
			if role == userRole {
				return true
			}
		}
	}
	return false
}

type DynamicRoutePermission struct {
	Pattern *regexp.Regexp
	Methods map[string][]models.Role
}

var dynamicPermissions = []DynamicRoutePermission{
	{
		Pattern: regexp.MustCompile(`^/api/users/\d+$`),
		Methods: map[string][]models.Role{
			"GET":    {models.Gerente, models.Admin, models.Caixa},
			"PUT":    {models.Gerente, models.Admin},
			"DELETE": {models.Gerente, models.Admin},
		},
	},
	{
		Pattern: regexp.MustCompile(`^/api/menuitem/\d+$`),
		Methods: map[string][]models.Role{
			"GET":    {models.Gerente, models.Admin, models.Cozinha},
			"PUT":    {models.Gerente, models.Admin, models.Cozinha},
			"DELETE": {models.Gerente, models.Admin},
		},
	},
	{
		Pattern: regexp.MustCompile(`^/api/menuitem/name/.*$`),
		Methods: map[string][]models.Role{
			"GET": {models.Cliente, models.Garcom, models.Gerente, models.Admin, models.Cozinha},
		},
	},
	{
		Pattern: regexp.MustCompile(`^/api/order/\d+$`),
		Methods: map[string][]models.Role{
			"GET":    {models.Cliente, models.Garcom, models.Gerente, models.Admin, models.Cozinha},
			"PUT":    {models.Garcom, models.Gerente, models.Admin, models.Cozinha},
			"DELETE": {models.Garcom, models.Gerente, models.Admin},
		},
	},
	{
		Pattern: regexp.MustCompile(`/api/order/user/\d+$`),
		Methods: map[string][]models.Role{
			"GET": {models.Cliente, models.Garcom, models.Gerente, models.Admin, models.Cozinha},
		},
	},
	{
		Pattern: regexp.MustCompile(`^/api/itemorder/\d+$`),
		Methods: map[string][]models.Role{
			"GET":    {models.Cliente, models.Garcom, models.Gerente, models.Admin, models.Cozinha},
			"PUT":    {models.Garcom, models.Gerente, models.Admin, models.Cozinha},
			"DELETE": {models.Garcom, models.Gerente, models.Admin},
		},
	},
	{
		Pattern: regexp.MustCompile(`^/api/users/password/\d+$`),
		Methods: map[string][]models.Role{
			"PUT": {models.Gerente, models.Admin},
		},
	},
}

// isAuthorized checks if the user is allowed to access the resource based on their role
func isAuthorized(userRole models.Role, path, method string) bool {
	// Check statics permissions
	if roles, ok := staticsPermissions[path][method]; ok {
		for _, role := range roles {
			if role == userRole {
				return true
			}
		}
	}

	// Check dynamics permissions
	for _, perm := range dynamicPermissions {
		if perm.Pattern.MatchString(path) {
			if roles, ok := perm.Methods[method]; ok {
				for _, role := range roles {
					if role == userRole {
						return true
					}
				}
			}
		}
	}

	log.Printf("Unauthorized userRole: %v  path: %s  method: %s", userRole, path, method)
	return false
}
