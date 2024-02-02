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

// UserService fornece métodos para operações relacionadas a usuários.
type UserService struct {
	store store.UsuarioStorer
}

// NewUsuarioService cria uma nova instância de UsuarioService.
func NewUsuarioService(store store.UsuarioStorer) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser trata da lógica de negócios para criar um novo usuário.
func (service *UserService) CreateUser(u *models.User) error {
	// Validar e-mail.
	if err := validarEmail(u.Email); err != nil {
		log.Print("erro CreateUsuario: ", err)
		return fmt.Errorf("erro CreateUsuario: %v", err)
	}
	// Validar papel
	if err := validarPapel(u.Email, string(u.Papel)); err != nil {
		log.Print("Validar Papel: ", err)
		return err
	}

	// Verificar se o e-mail já está em uso:
	existingUser, err := service.store.GetUserByEmail(u.Email)
	if err != nil {
		// Verifica se o erro é um erro de "nenhum registro encontrado"
		if err == sql.ErrNoRows {
			// Não é realmente um erro neste caso, então continue.
			log.Print("Usuário não encontrado, pronto para criar um novo.")
		} else {
			log.Print("service.store.GetUsuarioByEmail: ", err)
			return err
		}
	}
	if existingUser != nil {
		return fmt.Errorf("email '%s' já está em uso", u.Email)
	}

	// Continuar com a criação do usuário
	return service.store.CreateUser(u)
}

// GetUser trata da lógica para recuperar um usuário pelo ID.
func (service *UserService) GetUser(id int) (*models.User, error) {
	return service.store.GetUser(id)
}

// UpdateUser trata da lógica para atualizar um usuário existente.
func (service *UserService) UpdateUser(u *models.User) error {
	return service.store.UpdateUser(u)
}

// DeleteUser trata da lógica para deletar um usuário.
func (service *UserService) DeleteUser(id int) error {
	return service.store.DeleteUser(id)
}

// validarEmail valida o email
func validarEmail(email string) error {
	strRE := `^[a-zA-Z0-9.%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if ok, _ := regexp.MatchString(strRE, email); !ok {
		return errors.New("endereço de e-mail inválido")
	}
	return nil
}

// validarPapel verifica se o email iniciar com mesa[0-9]{2,}, número da mesa.
// Neste caso o papel só pode ser cliente
func validarPapel(email string, papel string) error {
	strRE := `^mesa[0-9]{2,}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if ok, _ := regexp.MatchString(strRE, email); ok {
		if papel != string(models.Cliente) {
			return errors.New("e-mail mesaXX@... devem ter papel == cliente")
		}
	}

	return nil
}
