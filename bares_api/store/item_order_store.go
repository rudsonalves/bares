package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"
)

const (
	createItemOrderSQL = "INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)"
	getItemOrderSQL    = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
	updateItemOrderSQL = "UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?"
	deleteItemOrderSQL = "DELETE FROM %s WHERE %s = ?"
)

// ItemOrderStore mantém a conexão com o banco de dados para operações
// relacionadas a itens pedidos.
type ItemOrderStore struct {
	DB *sql.DB
}

// NewItemOrder cria uma nova instância de ItemPedidoStore.
func NewItemOrder(db *sql.DB) *ItemOrderStore {
	return &ItemOrderStore{DB: db}
}

// ItemPedidoStoreStorer define as operações que um ItemPedidoStoreStore
// precisa implementar.
type ItemOrderStorer interface {
	CreateItemOrder(item *models.ItemOrder) error
	GetItemOrder(id int) (*models.ItemOrder, error)
	UpdateItemOrder(item *models.ItemOrder) error
	DeleteItemOrder(id int) error
}

// Garanta que ItemPedidoStoreStore implementa ItemPedidoStoreStorer.
var _ ItemOrderStorer = &ItemOrderStore{}

// CreateItemOrder adiciona um novo ItemPedido ao banco de dados.
func (store *ItemOrderStore) CreateItemOrder(item *models.ItemOrder) error {
	sqlString := fmt.Sprintf(createItemOrderSQL, TableItensOrders, OrderId, ItemId, Amount, Comments)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro CreateItemPedido: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.OrderId, item.ItemId, item.Amount, item.Comments)
	if err != nil {
		log.Printf("erro CreateItemPedido: %v", err)
		return err
	}
	itemPedidoID, err := result.LastInsertId()
	if err != nil {
		log.Printf("erro CreateItemPedido: %v", err)
		return err
	}
	item.Id = int(itemPedidoID)

	return nil
}

// GetItemOrder busca um itemPedido pelo ID.
func (store *ItemOrderStore) GetItemOrder(id int) (*models.ItemOrder, error) {
	item := &models.ItemOrder{}

	sqlString := fmt.Sprintf(getItemOrderSQL, Id, OrderId, ItemId, Amount,
		Comments, TableItensOrders, Id)

	err := store.DB.QueryRow(sqlString, id).Scan(
		&item.Id,
		&item.OrderId,
		&item.ItemId,
		&item.Amount,
		&item.Comments,
	)
	if err != nil {
		log.Printf("erro GetItemPedido: %v", err)
		return nil, err
	}

	return item, nil
}

// UpdateItemOrder atualiza os dados de um itemPedido.
func (store *ItemOrderStore) UpdateItemOrder(item *models.ItemOrder) error {
	sqlString := fmt.Sprintf(updateItemOrderSQL, TableItensOrders, OrderId, ItemId,
		Amount, Comments, Id)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdateItemPedido: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.OrderId, item.ItemId, item.Amount, item.Comments, item.Id)
	if err != nil {
		log.Printf("erro UpdateItemPedido: %v", err)
		return err
	}

	return nil
}

// DeleteItemOrder remove um itemPedido do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *ItemOrderStore) DeleteItemOrder(id int) error {
	sqlString := fmt.Sprintf(deleteItemOrderSQL, TableItensOrders, Id)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro DeleteItemPedido: %v", err)
		return fmt.Errorf("erro DeleteItemPedido: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("erro DeleteItemPedido: %v", err)
		return fmt.Errorf("erro DeleteItemPedido: %v", err)
	}

	return nil
}
