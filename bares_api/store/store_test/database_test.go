package storetest_test

import (
	"bares_api/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDatabase(t *testing.T) {
  // Cria um mock do banco de dados e do sqlmock
  db, mock, err := sqlmock.New()
  if err != nil {
    t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
  }
  defer db.Close()

  sqlExpected := []string{
    `CREATE DATABASE IF NOT EXISTS BarDB_test`,
    `USE BarDB_test`,
    `CREATE TABLE IF NOT EXISTS Usuarios \(
      usuarioID INT AUTO_INCREMENT PRIMARY KEY,
      nome VARCHAR\(255\) NOT NULL,
      email VARCHAR\(255\) UNIQUE NOT NULL,
      senhaHash VARCHAR\(255\) NOT NULL,
      papel ENUM\('cliente', 'garcom', 'gerente'\) NOT NULL
    \)`,
    `CREATE TABLE IF NOT EXISTS ItensMenu \(
          itemID INT AUTO_INCREMENT PRIMARY KEY,
          nome VARCHAR\(255\) UNIQUE NOT NULL,
          descricao TEXT,
          preco DECIMAL\(10,2\) NOT NULL,
          imagemURL VARCHAR\(255\)
        \)`,
    `CREATE INDEX IF NOT EXISTS idx_itensMenu_name ON ItensMenu\(nome\)`,
    `CREATE TABLE IF NOT EXISTS Pedidos \(
      pedidoID INT AUTO_INCREMENT PRIMARY KEY,
      usuarioID INT NOT NULL,
      dataHora DATETIME NOT NULL,
      status ENUM\('recebido', 'preparando', 'pronto', 'entregue'\) NOT NULL,
      FOREIGN KEY \(usuarioID\) REFERENCES Usuarios\(usuarioID\)
    \)`,
    `CREATE INDEX IF NOT EXISTS idx_usuario_id ON Pedidos\(usuarioID\)`,
    `CREATE TABLE IF NOT EXISTS ItensPedido \(
      itemPedidoID INT AUTO_INCREMENT PRIMARY KEY,
      pedidoID INT NOT NULL,
      itemID INT NOT NULL,
      quantidade INT DEFAULT 1,
      observacoes VARCHAR\(255\),
      FOREIGN KEY \(pedidoID\) REFERENCES Pedidos\(pedidoID\),
      FOREIGN KEY \(itemID\) REFERENCES ItensMenu\(itemID\)
    \)`,
    `CREATE INDEX IF NOT EXISTS idx_pedido_id ON ItensPedido\(pedidoID\)`,
  }

  // Prepara o mock para esperar as chamadas que serão feitas pela função CreateDatabase
  for _, sqlString := range sqlExpected {
    mock.ExpectExec(sqlString).WillReturnResult(sqlmock.NewResult(1, 1))
  }

  store := store.NewDatabaseStore("BarDB_test", db)

  // Chama a função CreateDatabase
  err = store.CreateDatabase()
  assert.NoError(t, err)

  // Verifica se todas as expectativas foram atendidas
  if err := mock.ExpectationsWereMet(); err != nil {
    t.Errorf("there were unfulfilled expectations: %s", err)
  }
}
