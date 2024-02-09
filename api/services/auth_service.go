package services

import (
	"bares_api/models"
	"bares_api/store"

	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// AuthService manages operations related to user authentication.
type AuthService struct {
	UsuarioStore *store.UserStore
}

// NewAuthService creates a new instance of the AuthService service.
func NewAuthservice(usuarioService *store.UserStore) *AuthService {
	return &AuthService{
		UsuarioStore: usuarioService,
	}
}

// ValidateCredentials checks the provided credentials and returns the username and a possible error.
func (service *AuthService) ValidateCredentials(credentials models.Credentials) (*models.User, error) {
	// Try to search for the user by their email ID.
	user, err := service.UsuarioStore.GetUserByEmail(credentials.Email)
	if err != nil {
		// Returns an error if the user is not found.
		return nil, fmt.Errorf("usuário %s não encontrado: %s", credentials.Email, err)
	}

	// Compares the provided password with the stored password hash.
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password))
	if err != nil {
		// Returns an error if the password is incorrect
		return nil, fmt.Errorf("senha do usuário %s incorreta: %s", credentials.Email, err)
	}

	// Returns the user and nil for error if the password is correct.
	return user, nil
}
