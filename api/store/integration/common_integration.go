package integration

import (
	"bares_api/models"
	"bares_api/store"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Configurar conexão com o banco de dados de teste
const testDBName = "BarDB_test"

// CreateUsers cria e insere uma lista de usuários de teste no banco de dados usando o UsuarioStore fornecido.
// Retorna a lista de usuários criados e um erro, se ocorrer.
func CreateUsers(s *store.UserStore) ([]models.User, error) {
	// Criar Usuários
	users := []models.User{
		{
			Name:         "John Doe",
			Email:        "mesa01@example.com",
			PasswordHash: "hashedpassword",
			Role:         "cliente",
		},
		{
			Name:         "Sergio Sacani",
			Email:        "masa02@example.com",
			PasswordHash: "1111111",
			Role:         "cliente",
		},
		{
			Name:         "Solange Almeida",
			Email:        "mesa03@email.com",
			PasswordHash: "22222222",
			Role:         "cliente",
		},
		{
			Name:         "Eduarda Carneiro",
			Email:        "eduarda@email.com",
			PasswordHash: "12341234",
			Role:         "gerente",
		},
		{
			Name:         "Gilberto Soares",
			Email:        "soares@email.com",
			PasswordHash: "qwerqwer",
			Role:         "garcom",
		},
	}

	// Inserir usuários no banco de dados
	for i, user := range users {
		err := s.CreateUser(&user)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar usuário: %s", err)
		}
		if user.Id == 0 {
			return nil, fmt.Errorf("id do usuário não foi definido após a criação")
		}
		// Atualizar a lista de usuários com IDs atribuídos
		users[i].Id = user.Id
	}

	return users, nil
}

// CreateItensMenu cria e insere uma lista de itemMenus de teste no banco de dados usando o ItemMenuStore fornecido.
// Retorna a lista de itemMenus criados e um erro, se ocorrer.
func CreateItensMenu(s *store.MenuItemStore) ([]models.MenuItem, error) {
	// Criar itemMenus
	itens := []models.MenuItem{
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

	// Inserir itemMenus no banco de dados
	for index, item := range itens {
		err := s.CreateMenuItem(&item)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar itemMenu: %s", err)
		}
		if item.Id == 0 {
			return nil, fmt.Errorf("id do itemMenu não foi definido após a criação")
		}
		// Atualizar a lista de itemMenus com IDs atribuídos
		itens[index].Id = item.Id
	}

	return itens, nil
}

// CreateItensMenu cria e insere uma lista de itemMenus de teste no banco de dados usando o ItemMenuStore fornecido.
// Retorna a lista de itemMenus criados e um erro, se ocorrer.
func CreateOrders(s *store.OrderStore) ([]models.Order, error) {
	// Criar Pedidos
	pedidos := []models.Order{
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

	// Inserir itemMenus no banco de dados
	for i, pedido := range pedidos {
		err := s.CreateOrder(&pedido)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar itemMenu: %s", err)
		}
		if pedido.Id == 0 {
			return nil, fmt.Errorf("id do itemMenu não foi definido após a criação")
		}
		// Atualizar a lista de itemMenus com IDs atribuídos
		pedidos[i].Id = pedido.Id
	}

	return pedidos, nil
}

// StartDatabase inicia o banco de dados
func StartDatabase() (*store.DatabaseStore, error) {
	log.Printf("Banco de dados teste: %q", testDBName)

	// Abre a conexão com o banco de dados
	connString := "alves_test:1234qwer@tcp(localhost:3306)/"
	dbStore, err := store.DatabaseOpen(testDBName, connString)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco de dados de teste: %s", err)
	}
	// defer dbStore.DatabaseClose() // Garante que a conexão será fechada no final

	// Limpeza inicial
	if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName)); err != nil {
		fmt.Printf("Falha ao limpar o banco de dados de teste: %s\n", err)
	}

	// CreateDatabase
	err = dbStore.CreateDatabase()
	if err != nil {
		return nil, fmt.Errorf("falha ao criar banco de dados e tabelas: %s", err)
	}

	return dbStore, nil
}
