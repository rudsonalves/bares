package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"
)


const (
  createItemOrderSQL = "INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)"
  getItemOrderSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
  updateItemOrderSQL = "UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?"
  deleteItemOrderSQL = "DELETE FROM %s WHERE %s = ?"
)

// ItemPedidoStore mantém a conexão com o banco de dados para operações
// relacionadas a itens pedidos.
type ItemPedidoStore struct {
  DB *sql.DB
}

// NewItemPedido cria uma nova instância de ItemPedidoStore.
func NewItemPedido(db *sql.DB) *ItemPedidoStore {
  return &ItemPedidoStore{DB: db}
}

// ItemPedidoStoreStorer define as operações que um ItemPedidoStoreStore
// precisa implementar.
type ItemPedidoStorer interface {
  CreateItemPedido(item *models.ItemPedido) error
  GetItemPedido(id int) (*models.ItemPedido, error)
  UpdateItemPedido(item *models.ItemPedido) error
  DeleteItemPedido(id int) error
}

// Garanta que ItemPedidoStoreStore implementa ItemPedidoStoreStorer.
var _ ItemPedidoStorer = &ItemPedidoStore{}

// CreateItemPedido adiciona um novo ItemPedido ao banco de dados.
func (store *ItemPedidoStore) CreateItemPedido(item *models.ItemPedido) error {
  sqlString := fmt.Sprintf(createItemOrderSQL, TableItensPedido, PedidoID, ItemID, Quantidade, Observacoes)

  stmt, err := store.DB.Prepare(sqlString)
  if err != nil {
    log.Printf("erro CreateItemPedido: %v", err)
		return fmt.Errorf("erro CreateItemPedido: %v", err)
  }
  defer stmt.Close()

  result, err := stmt.Exec(item.PedidoID, item.ItemID, item.Quantidade, item.Observacoes)
  if err != nil {
    log.Printf("erro CreateItemPedido: %v", err)
		return fmt.Errorf("erro CreateItemPedido: %v", err)
  }
  itemPedidoID, err := result.LastInsertId()
  if err != nil {
    log.Printf("erro CreateItemPedido: %v", err)
		return fmt.Errorf("erro CreateItemPedido: %v", err)
  }
  item.ItemPedidoID = int(itemPedidoID)

  return nil
}

// GetItemPedido busca um itemPedido pelo ID.
func (store *ItemPedidoStore) GetItemPedido(id int) (*models.ItemPedido, error) {
  item := &models.ItemPedido{}

  sqlString := fmt.Sprintf(getItemOrderSQL, ItemPedidoID, PedidoID, ItemID, Quantidade,
    Observacoes, TableItensPedido, ItemPedidoID)

  err := store.DB.QueryRow(sqlString, id).Scan(
    &item.ItemPedidoID,
    &item.PedidoID,
    &item.ItemID,
    &item.Quantidade,
    &item.Observacoes,
  )
  if err != nil {
    log.Printf("erro GetItemPedido: %v", err)
		return nil, fmt.Errorf("erro GetItemPedido: %v", err)
  }

  return item, nil
}

// UpdateItemPedido atualiza os dados de um itemPedido.
func (store *ItemPedidoStore) UpdateItemPedido(item *models.ItemPedido) error {
  sqlString := fmt.Sprintf(updateItemOrderSQL, TableItensPedido, PedidoID, ItemID,
    Quantidade, Observacoes, ItemPedidoID)

  stmt, err := store.DB.Prepare(sqlString)
  if err != nil {
    log.Printf("erro UpdateItemPedido: %v", err)
		return fmt.Errorf("erro UpdateItemPedido: %v", err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(item.PedidoID, item.ItemID, item.Quantidade, item.Observacoes, item.ItemPedidoID)
  if err != nil {
    log.Printf("erro UpdateItemPedido: %v", err)
		return fmt.Errorf("erro UpdateItemPedido: %v", err)
  }

  return nil
}

// DeleteItemPedido remove um itemPedido do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *ItemPedidoStore) DeleteItemPedido(id int) error {
  sqlString := fmt.Sprintf(deleteItemOrderSQL, TableItensPedido, ItemPedidoID)

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
