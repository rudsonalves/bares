package services

import (
	"bares_api/models"
	"bares_api/store"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
)

// UserService provides methods for user-related operations.
type UserService struct {
	store store.UsuarioStorer
}

type UserServiceInterface interface {
	CreateUser(u *models.User) error
	GetUser(id int) (*models.User, error)
	UpdateUser(u *models.User) error
	DeleteUser(id int) error
	CheckIfAdminExists() (bool, error)
}

// Ensure that UserService implements User Service Interface.
var _ UserServiceInterface = &UserService{}

// NewUsuarioService creates a new instance of UsuarioService.
func NewUsuarioService(store store.UsuarioStorer) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser handles the business logic for creating a new user.
func (service *UserService) CreateUser(u *models.User) error {
	// Validate e-mail.
	if err := validarEmail(u.Email); err != nil {
		log.Print("erro CreateUsuario: ", err)
		return fmt.Errorf("erro CreateUsuario: %v", err)
	}
	// Validate papel
	if err := validateRope(u.Email, string(u.Role)); err != nil {
		log.Print("Validar Papel: ", err)
		return err
	}

	// Check if the email is already in use
	existingUser, err := service.store.GetUserByEmail(u.Email)
	if err != nil {
		// Checks if the error is a "no records found" error
		if err == sql.ErrNoRows {
			// It's not really an error in this case, so continue.
			log.Print("Usuário não encontrado, pronto para criar um novo.")
		} else {
			log.Print("service.store.GetUsuarioByEmail: ", err)
			return err
		}
	}
	if existingUser != nil {
		return fmt.Errorf("email '%s' já está em uso", u.Email)
	}

	//Continue with user creation
	return service.store.CreateUser(u)
}

// GetUser handles the logic to retrieve a user by ID.
func (service *UserService) GetUser(id int) (*models.User, error) {
	return service.store.GetUser(id)
}

// UpdateUser handles the logic for update an existing user.
func (service *UserService) UpdateUser(u *models.User) error {
	return service.store.UpdateUser(u)
}

// DeleteUser handles the logic for delete an user.
func (service *UserService) DeleteUser(id int) error {
	return service.store.DeleteUser(id)
}

// CheckIfAdminExists checks if an administration user exists in the database
func (service *UserService) CheckIfAdminExists() (bool, error) {
	users, err := service.store.GetUsersByRole(models.Admin)
	if err != nil {
		return false, err
	}
	if len(users) > 0 {
		return true, nil
	}
	return false, nil
}

// validarEmail validate the email
func validarEmail(email string) error {
	strRE := `^[a-zA-Z0-9.%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if ok, _ := regexp.MatchString(strRE, email); !ok {
		return errors.New("endereço de e-mail inválido")
	}
	return nil
}

// validateRope checks if the email starts with mesa[0-9]{2,}, 'mesa' + table number.
// In this case the role can only be 'cliente'.
func validateRope(email string, papel string) error {
	strRE := `^mesa[0-9]{2,}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if ok, _ := regexp.MatchString(strRE, email); ok {
		if papel != string(models.Cliente) {
			return errors.New("e-mail mesaXX@... devem ter papel == cliente")
		}
	}

	return nil
}
