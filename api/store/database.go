package store

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	TableUsers   = "UsersTable"
	Id           = "id"
	Name         = "name"
	Email        = "email"
	PasswordHash = "passwordHash"
	Role         = "role"

	TableMenuItem = "MenuItemsTable"
	// Id        = "id"
	// Name           = "name"
	Description = "description"
	Price       = "price"
	ImageURL    = "imageURL"

	IndexMenuItem = "idx_menuItems_name"

	TableOrders = "OrdersTable"
	// Id     = "id"
	UserId   = "userId"
	DateHour = "dateHour"
	Status   = "status"

	IndexUserId = "idx_users_id"

	TableItensOrders = "ItemsOrderTable"
	// Id      = "id"
	OrderId  = "orderId"
	ItemId   = "itemId"
	Amount   = "amount"
	Comments = "comments"

	IndexOrderId = "idx_order_id"
)

// DatabaseStore mantém a conexão com o banco de dados.
type DatabaseStore struct {
	DBName string
	DB     *sql.DB
}

// NewDatabaseStore cria uma nova instância de DatabaseStore.
func NewDatabaseStore(dbName string, db *sql.DB) *DatabaseStore {
	return &DatabaseStore{
		DBName: dbName,
		DB:     db,
	}
}

// DatabaseOpen abre a conexão com o banco de dados.
func DatabaseOpen(dbName string, connString string) (*DatabaseStore, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	// Testa a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	store := NewDatabaseStore(dbName, db)
	return store, nil
}

// CreateDatabase cria o banco de dados e todas as suas tabelas, se necessário.
func (store *DatabaseStore) CreateDatabase() error {
	if err := store.createDatabaseIfNotExists(); err != nil {
		return err
	}

	if err := store.useDatabase(); err != nil {
		return err
	}

	if err := store.createTables(); err != nil {
		return err
	}

	if err := store.createIndexes(); err != nil {
		return err
	}

	return nil
}

// createDatabaseIfNotExists cria o banco de dados, caso este não exista
func (store *DatabaseStore) createDatabaseIfNotExists() error {
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", store.DBName)
	_, err := store.DB.Exec(createDBSQL)
	if err != nil {
		log.Printf("erro ao criar o banco de dados: %v", err)
		return err
	}
	return nil
}

// useDatabase usa o banco de dados store.DBName
func (store *DatabaseStore) useDatabase() error {
	useDBSQL := fmt.Sprintf("USE %s", store.DBName)
	_, err := store.DB.Exec(useDBSQL)
	if err != nil {
		log.Printf("erro ao selecionar o banco de dados: %v", err)
		return err
	}
	return nil
}

// createTables cria as tabelas TableUsuarios, TableItensMenu, TablePedidos e
// TableItensPedido no banco de dados
func (store *DatabaseStore) createTables() error {
	createTableSQLs := []string{
		fmt.Sprintf(`
      CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s VARCHAR(255) NOT NULL,
        %s VARCHAR(255) UNIQUE NOT NULL,
        %s VARCHAR(255) NOT NULL,
        %s ENUM('cliente', 'garcom', 'gerente', 'admin', 'cozinha', 'caixa') NOT NULL
      )`, TableUsers, Id, Name, Email, PasswordHash, Role,
		),
		fmt.Sprintf(`
      CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s VARCHAR(255) UNIQUE NOT NULL,
        %s TEXT,
        %s DECIMAL(10,2) NOT NULL,
        %s VARCHAR(255)
      )`, TableMenuItem, Id, Name, Description, Price, ImageURL,
		),
		fmt.Sprintf(`
      CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s INT NOT NULL,
        %s DATETIME NOT NULL,
        %s ENUM('recebido', 'preparando', 'pronto', 'entregue') NOT NULL,
        FOREIGN KEY (%s) REFERENCES %s(%s)
      )`, TableOrders, Id, UserId, DateHour, Status, UserId, TableUsers, Id,
		),
		fmt.Sprintf(`
      CREATE TABLE IF NOT EXISTS %s (
          %s INT AUTO_INCREMENT PRIMARY KEY,
        %s INT NOT NULL,
        %s INT NOT NULL,
        %s INT DEFAULT 1,
        %s VARCHAR(255),
        FOREIGN KEY (%s) REFERENCES %s(%s),
        FOREIGN KEY (%s) REFERENCES %s(%s)
      )`, TableItensOrders, Id, OrderId, ItemId, Amount, Comments,
			OrderId, TableOrders, Id,
			ItemId, TableMenuItem, Id,
		),
	}

	for _, createTableSQL := range createTableSQLs {
		stmt, err := store.DB.Prepare(createTableSQL)
		if err != nil {
			log.Printf("createTables: %v", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			log.Printf("createTables: %v", err)
			return err
		}
	}

	return nil
}

// createIndexes cria os índices das tabelas
func (store *DatabaseStore) createIndexes() error {
	createIndexSQLs := []string{
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", IndexMenuItem, TableMenuItem, Name),
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", IndexUserId, TableOrders, UserId),
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", IndexOrderId, TableItensOrders, OrderId),
	}

	for _, createIndexSQL := range createIndexSQLs {
		stmt, err := store.DB.Prepare(createIndexSQL)
		if err != nil {
			log.Printf("createIndexes: %v", err)
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			log.Printf("createIndexes: %v", err)
			return err
		}
	}

	return nil
}

// DatabaseClose fecha a conexão com o banco de dados.
func (store *DatabaseStore) DatabaseClose() {
	err := store.DB.Close()
	if err != nil {
		log.Printf("DatabaseClose :%v", err)
	}
}
