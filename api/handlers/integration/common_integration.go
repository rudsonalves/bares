package integration

import (
	"bares_api/models"
	"bares_api/store"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// UsersGenerate cria e insere uma lista de usuários de teste no banco de dados usando o UsuarioStore fornecido.
// Retorna a lista de usuários criados e um erro, se ocorrer.
func UsersGenerate() []models.User {
	// Criar Usuários
	users := []models.User{
		{
			Name:         "John Doe",
			Email:        "mesa01@example.com",
			PasswordHash: "hashedpassword",
			Role:         models.Cliente,
		},
		{
			Name:         "Sergio Sacani",
			Email:        "masa02@example.com",
			PasswordHash: "1111111",
			Role:         models.Cliente,
		},
		{
			Name:         "Solange Almeida",
			Email:        "mesa03@email.com",
			PasswordHash: "22222222",
			Role:         models.Cliente,
		},
		{
			Name:         "Eduarda Carneiro",
			Email:        "eduarda@email.com",
			PasswordHash: "12341234",
			Role:         models.Gerente,
		},
		{
			Name:         "Gilberto Soares",
			Email:        "soares@email.com",
			PasswordHash: "qwerqwer",
			Role:         models.Garcom,
		},
	}

	return users
}

// MenuItemsGenerate cria e insere uma lista de itemMenus de teste no banco de dados usando o ItemMenuStore fornecido.
// Retorna a lista de itemMenus criados e um erro, se ocorrer.
func MenuItemsGenerate(s *store.MenuItemStore) []models.MenuItem {
	// Criar itemMenus
	items := []models.MenuItem{
		{
			Name:        "Salada de Frutas",
			Description: "Uma bela salada de muitas frutas",
			Price:       25.99,
			ImageURL:    "image/fig01.jpg",
		},
		{
			Name:        "Bife a Rolê",
			Description: "Bife de boi enrolado com cenoura e bacon",
			Price:       48.99,
			ImageURL:    "image/fig02.jpg",
		},
		{
			Name:        "Feijão Tropeiro",
			Description: "Feijão, farinha, linguiça calabresa e muito bacon",
			Price:       34.99,
			ImageURL:    "image/fig03.jpg",
		},
		{
			Name:        "Suco de Alho",
			Description: "Alho batido com limão ciciliano",
			Price:       12.99,
			ImageURL:    "image/fig05.jpg",
		},
	}

	return items
}

// OrdersGenerate cria e insere uma lista de itemMenus de teste no banco de dados usando o ItemMenuStore fornecido.
// Retorna a lista de itemMenus criados e um erro, se ocorrer.
func OrdersGenerate(s *store.OrderStore) []models.Order {
	// Criar Pedidos
	orders := []models.Order{
		{
			UserId:   2,
			DateHour: time.Now(),
			Status:   models.Preparando,
		},
		{
			UserId:   1,
			DateHour: time.Now(),
			Status:   models.Recebido,
		},
		{
			UserId:   3,
			DateHour: time.Now(),
			Status:   models.Pronto,
		},
		{
			UserId:   1,
			DateHour: time.Now(),
			Status:   models.Recebido,
		},
	}

	return orders
}
