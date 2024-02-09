package utils

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"

	"golang.org/x/term"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz[]_,.;:!@#$%&*-*/+"

// GenerateRandomPassword gera uma senha aleatória com o comprimento especificado.
// A senha pode conter caracteres alfanuméricos e símbolos especiais.
// O parâmetro 'length' define o comprimento da senha desejada.
// Retorna a senha gerada e um erro, se ocorrer.
func GenerateRandomPassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("o comprimento deve ser maior que zero")
	}

	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

// CreateConnString cria a string de conexão para o banco de dados.
func CreateConnString(dbName string) string {
	reader := bufio.NewReader(os.Stdin)
	var connectionString string

	for {
		fmt.Println("===================================")
		fmt.Println("=            Bares API            =")
		fmt.Println("=          Version 1.0.0          =")
		fmt.Println("===================================")
		fmt.Printf("\nUsuário e senha para o acesso ao banco de dados %s.\n", dbName)

		// Nome do usuário dp banco de dados
		fmt.Printf("Entre com o nome: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler o nome, tente novamente.")
			continue
		}
		name = TrimSpaceLB(name)

		// Senha do usuário do bando de dados
		fmt.Printf("Entre com a senha: ")
		passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Erro ao ler a senha, tente novamente.")
			continue
		}
		password := string(passwordBytes)

		// Confirmação da senha
		fmt.Printf("\nConfirme a senha: ")
		checkPasswordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Erro ao ler a senha, tente novamente.")
			continue
		}
		checkPassword := string(checkPasswordBytes)

		if password != checkPassword {
			fmt.Printf("\nSenha errada. Entre novamente!\n\n")
			continue
		}

		// cria string de conexão
		connectionString = fmt.Sprintf("%s:%s@tcp(localhost:3306)/", name, password)

		// Última verificação
		fmt.Printf("\n\nString de conexão: '%s:%s@tcp(localhost:3306)/'\n\n",
			name, strings.Repeat("*", len(password)))
		fmt.Println("Confirme as informações (S/n)? ")
		ans, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a confirmação, tente novamente.")
			continue
		}
		ans = TrimSpaceLB(ans)

		if strings.EqualFold(ans, "n") {
			fmt.Println("Credenciais rejeitadas!")
			continue
		}
		break
	}
	println("Credenciais confirmadas!")

	return connectionString
}

func TrimSpaceLB(answer string) string {
	return strings.TrimSpace(strings.Trim(strings.TrimSpace(answer), "\n"))
}

// IPCheck retorna a lista de IPs conectados na máquina.
func IPCheck() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("\nIPCheck: %s", err)
	}

	var ipList []string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipList = append(ipList, ipNet.IP.String())
				fmt.Println("IPv4:", ipNet.IP.String())
			} else if ipNet.IP.To16() != nil {
				ipList = append(ipList, ipNet.IP.String())
				fmt.Println("IPv6:", ipNet.IP.String())
			}
		}
	}

	if len(ipList) > 0 {
		return ipList[0], nil
	}

	return "", fmt.Errorf("error: No connected network device found")
}
