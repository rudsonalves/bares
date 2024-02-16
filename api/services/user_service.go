package services

import (
	"bares_api/models"
	"bares_api/store"
	"bares_api/utils"
	"database/sql"
	"fmt"
	"log"
)

// UserService provides methods for user-related operations.
type UserService struct {
	store store.UserStorer
}

type UserServiceInterface interface {
	CreateUser(u *models.User) error
	GetUser(id int) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(u *models.User) error
	UpdateUserPass(userId int, password string) error
	DeleteUser(id int) error
	CheckIfAdminExists() (bool, error)
}

// Ensure that UserService implements User Service Interface.
var _ UserServiceInterface = &UserService{}

// NewUsuarioService creates a new instance of UsuarioService.
func NewUsuarioService(store store.UserStorer) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser handles the business logic for creating a new user.
func (service *UserService) CreateUser(u *models.User) error {
	// Validate e-mail.
	if err := utils.ValidateEmail(u.Email); err != nil {
		log.Print("erro CreateUsuario: ", err)
		return fmt.Errorf("erro CreateUsuario: %v", err)
	}
	// Validate role
	if err := utils.ValidateRope(u.Email, string(u.Role)); err != nil {
		log.Print("Validar Role: ", err)
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

func (service *UserService) GetAllUsers() ([]*models.User, error) {
	return service.store.GetAllUsers()
}

// UpdateUser handles the logic for update an existing user.
func (service *UserService) UpdateUser(u *models.User) error {
	return service.store.UpdateUser(u)
}

// UpdateUserPass takes care of updating the user's password.
func (service *UserService) UpdateUserPass(userId int, password string) error {
	passwordStrength := utils.EvaluatePasswordStrength(password)
	// FIXME: Verificar esta parte do código.
	if passwordStrength.Score < 6 {
		log.Printf("error UpdateUserPass: %s", passwordStrength.Feedback)
	}
	return service.store.UpdateUserPass(userId, password)
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
