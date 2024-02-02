package services

import (
	"bares_api/models"
	"bares_api/store"
	"bares_api/utils"
	"log"

	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// AuthService gerencia as operações relacionadas à autenticação de usuários.
type AuthService struct {
	UsuarioStore *store.UsuarioStore
}

// NewAuthService cria uma nova instância do serviço AuthService.
func NewAuthservice(usuarioService *store.UsuarioStore) *AuthService {
	return &AuthService{
		UsuarioStore: usuarioService,
	}
}

// ValidarCredenciais verifica as credenciais fornecidas e retorna o papel do usuário e um possível erro.
// Se as credenciais pertencerem a um cliente e este não estiver registrado, um novo usuário cliente
// será criado com uma senha aleatória, ou uma padrão em caso de erro no GenerateRandomPassword.
func (service *AuthService) ValidarCredenciais(credentials models.Credentials) (models.Papel, error) {
	erClient := `^mesa[0-9]{2}@.*$`
	if ok, _ := regexp.MatchString(erClient, credentials.Email); ok {
		// Verifica se a mesa está aberta (conta no banco de dados)
		user, err := service.UsuarioStore.GetUsuarioByEmail(credentials.Email)
		if err != nil {
			// Se não houver conta, cria uma nova
			password, err := utils.GenerateRandomPassword(12)
			if err != nil {
				log.Printf("Erro na geração da senha aleatória: %s\n Usando senha padrão.", err)
				password = "u4087qw0y78y@#$"
			}
			newUser := models.Usuario{
				Nome:      credentials.Nome,
				Email:     credentials.Email,
				SenhaHash: password,
				Papel:     models.Cliente,
			}

			err = service.UsuarioStore.CreateUsuario(&newUser)
			if err != nil {
				return "", err
			}
			return models.Cliente, nil
		}
		return user.Papel, nil
	}

	// Outros usuários
	user, err := service.UsuarioStore.GetUsuarioByEmail(credentials.Email)
	if err != nil {
		return "", fmt.Errorf("usuário %s não encontrado", credentials.Email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.SenhaHash), []byte(credentials.Password))
	if err != nil {
		return "", fmt.Errorf("senha do usuário %s incorreta", credentials.Email)
	}

	return user.Papel, nil
}
