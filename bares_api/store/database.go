package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

const (
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
func DatabaseOpen(dbName string) (*DatabaseStore, error) {
  connString := "alves_test:1234qwer@tcp(localhost:3306)/"
  if !strings.HasSuffix(dbName, "_test") {
    connString = createConnString()
  }

  db, err := sql.Open("mysql", connString)
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
    return fmt.Errorf("erro ao criar o banco de dados: %v", err)
  }
  return nil
}

// useDatabase usa o banco de dados store.DBName
func (store *DatabaseStore) useDatabase() error {
  useDBSQL := fmt.Sprintf("USE %s", store.DBName)
  _, err := store.DB.Exec(useDBSQL)
  if err != nil {
    log.Printf("erro ao selecionar o banco de dados: %v", err)
    return fmt.Errorf("erro ao selecionar o banco de dados: %v", err)
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
        %s ENUM('cliente', 'garcom', 'gerente') NOT NULL
      )`, TableUsuarios, UsuarioID, Nome, Email, SenhaHash, Papel,
    ),
    fmt.Sprintf(`
      CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s VARCHAR(255) UNIQUE NOT NULL,
        %s TEXT,
        %s DECIMAL(10,2) NOT NULL,
        %s VARCHAR(255)
      )`, TableItensMenu, ItemID, Nome, Descricao, Preco, ImagemURL,
    ),
    fmt.Sprintf(`
      CREATE TABLE IF NOT EXISTS %s (
        %s INT AUTO_INCREMENT PRIMARY KEY,
        %s INT NOT NULL,
        %s DATETIME NOT NULL,
        %s ENUM('recebido', 'preparando', 'pronto', 'entregue') NOT NULL,
        FOREIGN KEY (%s) REFERENCES %s(%s)
      )`, TablePedidos, PedidoID, UsuarioID, DataHora, Status, UsuarioID, TableUsuarios, UsuarioID,
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
      )`, TableItensPedido, ItemPedidoID, PedidoID, ItemID, Quantidade, Observacoes,
      PedidoID, TablePedidos, PedidoID, ItemID, TableItensMenu, ItemID,
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
    fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", IndexItensMenu, TableItensMenu, Nome),
    fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", IndexUsuarioId, TablePedidos, UsuarioID),
    fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", IndexPedidoId, TableItensPedido, PedidoID),
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
    log.Fatal(err)
  }
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
