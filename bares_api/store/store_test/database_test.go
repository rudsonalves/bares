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
		`CREATE TABLE IF NOT EXISTS UsersTable \(
      id INT AUTO_INCREMENT PRIMARY KEY,
      name VARCHAR\(255\) NOT NULL,
      email VARCHAR\(255\) UNIQUE NOT NULL,
      passwordHash VARCHAR\(255\) NOT NULL,
      role ENUM\('cliente', 'garcom', 'gerente', 'admin', 'cozinha', 'caixa'\) NOT NULL
    \)`,
		`CREATE TABLE IF NOT EXISTS MenuItemsTable \(
      id INT AUTO_INCREMENT PRIMARY KEY,
      name VARCHAR\(255\) UNIQUE NOT NULL,
      description TEXT,
      price DECIMAL\(10,2\) NOT NULL,
      imageURL VARCHAR\(255\)
    \)`,
		`CREATE TABLE IF NOT EXISTS OrdersTable \(
      id INT AUTO_INCREMENT PRIMARY KEY,
      userId INT NOT NULL,
      dateHour DATETIME NOT NULL,
      status ENUM\('recebido', 'preparando', 'pronto', 'entregue'\) NOT NULL,
      FOREIGN KEY \(userId\) REFERENCES UsersTable\(id\)
    \)`,
		`CREATE TABLE IF NOT EXISTS ItemsOrderTable \(
      id INT AUTO_INCREMENT PRIMARY KEY,
      orderId INT NOT NULL,
      itemId INT NOT NULL,
      amount INT DEFAULT 1,
      comments VARCHAR\(255\),
      FOREIGN KEY \(orderId\) REFERENCES OrdersTable\(id\),
      FOREIGN KEY \(itemId\) REFERENCES MenuItemsTable\(id\)
    \)`,
		`CREATE INDEX IF NOT EXISTS idx_menuItems_name ON MenuItemsTable\(name\)`,
		`CREATE INDEX IF NOT EXISTS idx_users_id ON OrdersTable\(userId\)`,
		`CREATE INDEX IF NOT EXISTS idx_order_id ON ItemsOrderTable\(orderId\)`,
	}

	// Prepara o mock para esperar as chamadas que serão feitas pela função CreateDatabase
	for index, sqlString := range sqlExpected {
		if index > 1 {
			mock.ExpectPrepare(sqlString)
		}
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
