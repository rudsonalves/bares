package bootstrap

import (
	"bares_api/models"
	"bares_api/services"
	"bares_api/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/term"
)

// CheckAndCreateAdminUser verifica se existe um usuário Admin. Caso não exista
// um usuário admin deve ser criado.
func CheckAndCreateAdminUser(userService *services.UserService) error {
	exists, err := userService.CheckIfAdminExists()
	if err != nil {
		log.Printf("Erro ao verificar a existência de administradores: %v", err)
		return err
	}

	if exists {
		fmt.Println("Existe usuário administrador no sistema.")
		return nil
	}

	// Criar usuário admin
	fmt.Printf("\n\n\n----------------------------------------------------------\n")
	fmt.Printf("\n\n\nNão foi encontrado nenhum administrador neste sistema.\n")
	fmt.Println("É necessário criar uma conta de Administrador para continuar.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("===================================")
		fmt.Println("=   Entre com as informações      =")
		fmt.Println("=       do usuário Admin          =")
		fmt.Println("===================================")

		var name string
		fmt.Printf("\nEntre com o NOME: ")
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler o nome, tente novamente.")
			continue
		}
		name = utils.TrimSpaceLB(name)

		var email string
		fmt.Printf("Entre com o EMAIL: ")
		email, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler o email, tente novamente.")
			continue
		}
		email = utils.TrimSpaceLB(email)

		var password string
		var checkPassword string
		for {
			fmt.Println("Utilize uma senha de mínimo 8 caracteres, com letras e números")
			fmt.Printf("Entre com uma senha: ")
			passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Erro ao ler a senha, tente novamente.")
				continue
			}
			password = string(passwordBytes)
			fmt.Println()

			fmt.Printf("Verificação da senha: ")
			passwordBytes, err = term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Erro ao ler a senha, tente novamente.")
				continue
			}
			checkPassword = string(passwordBytes)
			fmt.Println()

			if password != checkPassword {
				fmt.Printf("\nSenha devergentes! Tente novamente.\n\n")
				continue
			}

			if !isPasswordStrong(password) {
				fmt.Printf("\n\nSenha fraca! User uma senha forte.\n")
				continue
			}

			break
		}

		user := models.User{
			Name:         name,
			Email:        email,
			PasswordHash: password,
			Role:         models.Admin,
		}
		fmt.Printf("\n\nCriar Usuário:\n")
		fmt.Println(user)
		fmt.Println("Confirme a criação do Admin (S/n): ")

		var ans string
		ans, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a confirmação, tente novamente.")
			continue
		}
		ans = utils.TrimSpaceLB(ans)

		if strings.EqualFold(ans, "s") || len(ans) == 0 {
			err = userService.CreateUser(&user)
			if err != nil {
				log.Printf("Falha ao criar o usuário administrador: %v", err)
				return err
			}
			fmt.Println("Usuário administrador criado com sucesso!")
			break
		} else if strings.EqualFold(ans, "n") {
			fmt.Println("Criação do usuário administrador cancelada.")
			continue
		} else {
			fmt.Println("Resposta inválida. Por favor, responda com 'S' para sim ou 'N' para não.")
		}
	}

	return nil
}

// isPasswordStrong retorna verdadeiro se a senha for forte. No momento, uma senha
// é forte se tiver
func isPasswordStrong(password string) bool {
	// Verifica se a senha tem ao menos 8 caracteres
	if len(password) < 8 {
		return false
	}

	lowercase, _ := regexp.Compile("[a-z]") // Verifica letra minúscula
	uppercase, _ := regexp.Compile("[A-Z]") // Verifica letra maiúscula
	number, _ := regexp.Compile("[0-9]")    // Verifica número
	// special, _ := regexp.Compile("[^a-zA-Z0-9]") // Verifica caractere especial

	return lowercase.MatchString(password) && uppercase.MatchString(password) &&
		number.MatchString(password) // && special.MatchString(password)
}
