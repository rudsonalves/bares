package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

const (
	BarDB = "BarDB"

	TableUsuarios = "Usuarios"
	UsuarioID     = "usuarioID"
	Nome          = "nome"
	Email         = "email"
	SenhaHash     = "senhaHash"
	Papel         = "papel"

	TableItensMenu = "ItensMenu"
	ItemID         = "itemID"
	// Nome           = "nome"
	Descricao = "descricao"
	Preco     = "preco"
	ImagemURL = "imagemURL"

	IndexItensMenu = "idx_itensMenu_name"

	TablePedidos = "Pedidos"
	PedidoID     = "pedidoID"
	// UsuarioID    = "usuarioID"
	DataHora = "dataHora"
	Status   = "status"

	IndexUsuarioId = "idx_usuario_id"

	TableItensPedido = "ItensPedido"
	ItemPedidoID     = "itemPedidoID"
	// PedidoID         = "pedidoID"
	// ItemID           = "itemID"
	Quantidade  = "quantidade"
	Observacoes = "observacoes"

	IndexPedidoId = "idx_pedido_id"
)

// DatabaseStore mantém a conexão com o banco de dados.
type DatabaseStore struct {
	DB *sql.DB
}

// newDatabaseStore cria uma nova instância de DatabaseStore.
func newDatabaseStore(db *sql.DB) *DatabaseStore {
	return &DatabaseStore{DB: db}
}

// DatabaseOpen abre a conexão com o banco de dados.
func DatabaseOpen() (*DatabaseStore, error) {
	connString := createConnString()
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	store := newDatabaseStore(db)
	return store, nil
}

// createConnString cria a string de conexão para o banco de dados.
func createConnString() string {
	var name string

	fmt.Printf("Entre com o nome do usuário: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nEntre com a senha: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("\n%s", err)
	}

	password := string(passwordBytes)
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/", name, password)

	return connectionString
}

// CreateDatabase cria as tabelas do banco de dados
func (store *DatabaseStore) CreateDatabase() error {
	createTableSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", BarDB)
	_, err := store.DB.Exec(createTableSql)
	if err != nil {
		return err
	}

	createTableSql = fmt.Sprintf("USE %s", BarDB)
	_, err = store.DB.Exec(createTableSql)
	if err != nil {
		return err
	}

	createTableSql = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s VARCHAR(255) NOT NULL,
        %s VARCHAR(255) UNIQUE NOT NULL,
        %s VARCHAR(255) NOT NULL,
        %s ENUM('cliente', 'garcom', 'gerente') NOT NULL
	)`, TableUsuarios, UsuarioID, Nome, Email, SenhaHash, Papel)

	_, err = store.DB.Exec(createTableSql)
	if err != nil {
		return err
	}

	createTableSql = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s VARCHAR(255) UNIQUE NOT NULL,
        %s TEXT,
        %s DECIMAL(10,2) NOT NULL,
        %s VARCHAR(255)
    )`, TableItensMenu, ItemID, Nome, Descricao, Preco, ImagemURL)
	_, err = store.DB.Exec(createTableSql)
	if err != nil {
		return err
	}

	createIndexSql := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)",
		IndexItensMenu, TableItensMenu, Nome)
	_, err = store.DB.Exec(createIndexSql)
	if err != nil {
		return err
	}

	createTableSql = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		%s INT AUTO_INCREMENT PRIMARY KEY,
		%s INT NOT NULL,
		%s DATETIME NOT NULL,
		%s ENUM('recebido', 'preparando', 'pronto', 'entregue') NOT NULL,
		FOREIGN KEY (%s) REFERENCES %s(%s)
	)`, TablePedidos, PedidoID, UsuarioID, DataHora, Status, UsuarioID, TableUsuarios, UsuarioID)
	_, err = store.DB.Exec(createTableSql)
	if err != nil {
		return err
	}

	createIndexSql = fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)",
		IndexUsuarioId, TablePedidos, UsuarioID)
	_, err = store.DB.Exec(createIndexSql)
	if err != nil {
		return err
	}

	createTableSql = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
		%s INT AUTO_INCREMENT PRIMARY KEY,
		%s INT NOT NULL,
		%s INT NOT NULL,
		%s INT DEFAULT 1,
		%s VARCHAR(255),
		FOREIGN KEY (%s) REFERENCES %s(%s),
    	FOREIGN KEY (%s) REFERENCES %s(%s)
	)`, TableItensPedido, ItemPedidoID, PedidoID, ItemID, Quantidade, Observacoes, PedidoID, TablePedidos,
		PedidoID, ItemID, TableItensMenu, ItemID)
	_, err = store.DB.Exec(createTableSql)
	if err != nil {
		return err
	}

	createIndexSql = fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)",
		IndexPedidoId, TableItensPedido, PedidoID)
	_, err = store.DB.Exec(createIndexSql)
	if err != nil {
		return err
	}

	return nil
}

// DatabaseClose fecha a conexão com o banco de dados.
func (store *DatabaseStore) DatabaseClose() {
	err := store.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
