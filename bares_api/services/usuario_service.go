package services

import (
	"bares_api/models"
	"bares_api/store"
	"errors"
	"fmt"
	"regexp"
)

// UsuarioService fornece métodos para operações relacionadas a usuários.
type UsuarioService struct {
	store store.UsuarioStorer
}

// NewUsuarioService cria uma nova instância de UsuarioService.
func NewUsuarioService(store store.UsuarioStorer) *UsuarioService {
	return &UsuarioService{
		store: store,
	}
}

// CreateUsuario trata da lógica de negócios para criar um novo usuário.
func (service *UsuarioService) CreateUsuario(u *models.Usuario) error {
	// Validar e-mail.
	if err := validarEmail(u.Email); err != nil {
		return err
	}
	// Validar papel
	if err := validarPapel(u.Email, string(u.Papel)); err != nil {
		return err
	}

	// Verificar se o e-mail já está em uso:
	existingUser, err := service.store.GetUsuarioByEmail(u.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("email '%s' já está em uso", u.Email)
	}

	// Continuar com a criação do usuário
	return service.store.CreateUsuario(u)
}

// GetUsuario trata da lógica para recuperar um usuário pelo ID.
func (service *UsuarioService) GetUsuario(id int) (*models.Usuario, error) {
	return service.store.GetUsuario(id)
}

// UpdateUsuario trata da lógica para atualizar um usuário existente.
func (service *UsuarioService) UpdateUsuario(u *models.Usuario) error {
	return service.store.UpdateUsuario(u)
}

// DeleteUsuario trata da lógica para deletar um usuário.
func (service *UsuarioService) DeleteUsuario(id int) error {
	return service.store.DeleteUsuario(id)
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
