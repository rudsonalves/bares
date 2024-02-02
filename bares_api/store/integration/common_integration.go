package integration

import (
	"bares_api/models"
	"bares_api/store"
	"fmt"
	"log"
	"time"
)

// Configurar conexão com o banco de dados de teste
const testDBName = "BarDB_test"

// CreateUsers cria e insere uma lista de usuários de teste no banco de dados usando o UsuarioStore fornecido.
// Retorna a lista de usuários criados e um erro, se ocorrer.
func CreateUsers(s *store.UsuarioStore) ([]models.Usuario, error) {
	// Criar Usuários
	users := []models.Usuario{
		{
			Nome:      "John Doe",
			Email:     "mesa01@example.com",
			SenhaHash: "hashedpassword",
			Papel:     "cliente",
		},
		{
			Nome:      "Sergio Sacani",
			Email:     "masa02@example.com",
			SenhaHash: "1111111",
			Papel:     "cliente",
		},
		{
			Nome:      "Solange Almeida",
			Email:     "mesa03@email.com",
			SenhaHash: "22222222",
			Papel:     "cliente",
		},
		{
			Nome:      "Eduarda Carneiro",
			Email:     "eduarda@email.com",
			SenhaHash: "12341234",
			Papel:     "gerente",
		},
		{
			Nome:      "Gilberto Soares",
			Email:     "soares@email.com",
			SenhaHash: "qwerqwer",
			Papel:     "garcom",
		},
	}

	// Inserir usuários no banco de dados
	for i, user := range users {
		err := s.CreateUsuario(&user)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar usuário: %s", err)
		}
		if user.UsuarioID == 0 {
			return nil, fmt.Errorf("id do usuário não foi definido após a criação")
		}
		// Atualizar a lista de usuários com IDs atribuídos
		users[i].UsuarioID = user.UsuarioID
	}

	return users, nil
}

// CreateItensMenu cria e insere uma lista de itemMenus de teste no banco de dados usando o ItemMenuStore fornecido.
// Retorna a lista de itemMenus criados e um erro, se ocorrer.
func CreateItensMenu(s *store.ItensMenuStore) ([]models.ItemMenu, error) {
	// Criar itemMenus
	itens := []models.ItemMenu{
		{
			Nome:      "Salada de Frutas",
			Descricao: "Uma bela salada de muitas frutas",
			Preco:     25.99,
			ImagemURL: "image/fig01.jpg",
		},
		{
			Nome:      "Bife a Rolê",
			Descricao: "Bife de boi enrolado com cenoura e bacon",
			Preco:     48.99,
			ImagemURL: "image/fig02.jpg",
		},
		{
			Nome:      "Feijão Tropeiro",
			Descricao: "Feijão, farinha, linguiça calabresa e muito bacon",
			Preco:     34.99,
			ImagemURL: "image/fig03.jpg",
		},
		{
			Nome:      "Suco de Alho",
			Descricao: "Alho batido com limão ciciliano",
			Preco:     12.99,
			ImagemURL: "image/fig05.jpg",
		},
	}

	// Inserir itemMenus no banco de dados
	for i, item := range itens {
		err := s.CreateItemMenu(&item)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar itemMenu: %s", err)
		}
		if item.ItemID == 0 {
			return nil, fmt.Errorf("id do itemMenu não foi definido após a criação")
		}
		// Atualizar a lista de itemMenus com IDs atribuídos
		itens[i].ItemID = item.ItemID
	}

	return itens, nil
}

// CreateItensMenu cria e insere uma lista de itemMenus de teste no banco de dados usando o ItemMenuStore fornecido.
// Retorna a lista de itemMenus criados e um erro, se ocorrer.
func CreatePedidos(s *store.PedidoStore) ([]models.Pedido, error) {
	// Criar Pedidos
	pedidos := []models.Pedido{
		{
			UsuarioID: 2,
			DataHora:  time.Now(),
			Status:    models.Preparando,
		},
		{
			UsuarioID: 1,
			DataHora:  time.Now(),
			Status:    models.Recebido,
		},
		{
			UsuarioID: 3,
			DataHora:  time.Now(),
			Status:    models.Pronto,
		},
		{
			UsuarioID: 1,
			DataHora:  time.Now(),
			Status:    models.Recebido,
		},
	}

	// Inserir itemMenus no banco de dados
	for i, pedido := range pedidos {
		err := s.CreatePedido(&pedido)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar itemMenu: %s", err)
		}
		if pedido.PedidoID == 0 {
			return nil, fmt.Errorf("id do itemMenu não foi definido após a criação")
		}
		// Atualizar a lista de itemMenus com IDs atribuídos
		pedidos[i].PedidoID = pedido.PedidoID
	}

	return pedidos, nil
}

// StartDatabase inicia o banco de dados
func StartDatabase() (*store.DatabaseStore, error) {
	log.Printf("Banco de dados teste: %q", testDBName)

	// Abre a conexão com o banco de dados
	dbStore, err := store.DatabaseOpen(testDBName)
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
