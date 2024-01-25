package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
)

// ItemPedidoStore mantém a conexão com o banco de dados para operações
// relacionadas a itens pedidos.
type ItemPedidoStore struct {
	DB *sql.DB
}

// NewItemPedidoStore cria uma nova instância de ItemPedidoStore.
func NewItemPedidoStore(db *sql.DB) *ItemPedidoStore {
	return &ItemPedidoStore{DB: db}
}

// ItemPedidoStoreStorer define as operações que um ItemPedidoStoreStore
// precisa implementar.
type ItemPedidoStorer interface {
	CreateItemPedidoStore(item *models.ItemPedido) error
	GetItemPedidoStore(id int) (*models.ItemPedido, error)
	UpdateItemPedidoStore(item *models.ItemPedido) error
	DeleteItemPedidoStore(id int) error
}

// Garanta que ItemPedidoStoreStore implementa ItemPedidoStoreStorer.
var _ ItemPedidoStorer = &ItemPedidoStore{}

// CreateItemPedidoStore adiciona um novo ItemPedido ao banco de dados.
func (store *ItemPedidoStore) CreateItemPedidoStore(item *models.ItemPedido) error {
	createString := fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)",
		TableItensPedido, PedidoID, ItemID, Quantidade, Observacoes)

	stmt, err := store.DB.Prepare(createString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.PedidoID, item.ItemID, item.Quantidade, item.Observacoes)
	if err != nil {
		return err
	}
	itemPedidoID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	item.ItemPedidoID = int(itemPedidoID)

	return nil
}

// GetItemPedidoStore busca um itemPedido pelo ID.
func (store *ItemPedidoStore) GetItemPedidoStore(id int) (*models.ItemPedido, error) {
	i := &models.ItemPedido{}

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		ItemPedidoID, PedidoID, ItemID, Quantidade, Observacoes, TableItensPedido, ItemPedidoID)

	err := store.DB.QueryRow(queryString, id).Scan(&i.ItemPedidoID, &i.PedidoID, &i.ItemID, &i.Quantidade, &i.Observacoes)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// UpdateItemPedidoStore atualiza os dados de um itemPedido.
func (store *ItemPedidoStore) UpdateItemPedidoStore(item *models.ItemPedido) error {
	updateString := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?",
		TableItensPedido, PedidoID, ItemID, Quantidade, Observacoes, ItemPedidoID)

	stmt, err := store.DB.Prepare(updateString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.PedidoID, item.ItemID, item.Quantidade, item.Observacoes, item.ItemPedidoID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItemPedidoStore remove um itemPedido do banco de dados.
func (store *ItemPedidoStore) DeleteItemPedidoStore(id int) error {
	deleteString := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", TableItensPedido, ItemPedidoID)

	stmt, err := store.DB.Prepare(deleteString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
