package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"
)

const (
	createItemMenuSQL = "INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)"
	getItemMenuSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
	updateItemMenuSQL = "UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?"
	deleteItemMenuSQL = "DELETE FROM %s WHERE %s = ?"
	getALLItemMenuSQL = "SELECT %s, %s, %s, %s, %s FROM %s ORDER BY %s"
	getItemMenuByNameSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
)

// ItensMenuStore mantém a conexão com o banco de dados para operações relacionadas a itens do menu.
type ItensMenuStore struct {
	DB *sql.DB
}

// NewItensMenu cria uma nova instância de ItensMenuStorer.
func NewItensMenu(db *sql.DB) *ItensMenuStore {
	return &ItensMenuStore{DB: db}
}

// ItensMenuStorer define as operações que um ItensMenuStore precisa implementar.
type ItensMenuStorer interface {
	CreateItemMenu(item *models.ItemMenu) error
	GetItemMenu(id int) (*models.ItemMenu, error)
	UpdateItemMenu(item *models.ItemMenu) error
	DeleteItemMenu(id int) error
	GetAllItemMenu() ([]*models.ItemMenu, error)
	GetItemMenuByNome(nome string) (*models.ItemMenu, error)
}

// Garanta que ItensMenuStore implementa ItensMenuStorer.
var _ ItensMenuStorer = &ItensMenuStore{}

// CreateItemMenu adiciona um novo usuário ao banco de dados.
func (store *ItensMenuStore) CreateItemMenu(item *models.ItemMenu) error {
	sqlString := fmt.Sprintf(createItemMenuSQL,
		TableItensMenu, Nome, Descricao, Preco, ImagemURL)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro CreateItemMenu: %v", err)
		return fmt.Errorf("erro CreateItemMenu: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.Nome, item.Descricao, item.Preco, item.ImagemURL)
	if err != nil {
		log.Printf("erro CreateItemMenu: %v", err)
		return fmt.Errorf("erro CreateItemMenu: %v", err)
	}

	itemID, err := result.LastInsertId()
	if err != nil {
		log.Printf("erro CreateItemMenu: %v", err)
		return fmt.Errorf("erro CreateItemMenu: %v", err)
	}
	item.ItemID = int(itemID)

	return nil
}

// GetItemMenu busca um ItemMenu pelo ID.
func (store *ItensMenuStore) GetItemMenu(id int) (*models.ItemMenu, error) {
	i := &models.ItemMenu{}

	sqlString := fmt.Sprintf(getItemMenuSQL, ItemID, Nome, Descricao, Preco, ImagemURL,
		TableItensMenu, ItemID)

	err := store.DB.QueryRow(sqlString, id).Scan(&i.ItemID, &i.Nome, &i.Descricao, &i.Preco, &i.ImagemURL)
	if err != nil {
		log.Printf("erro GetItemMenu: %v", err)
		return nil, fmt.Errorf("erro GetItemMenu: %v", err)
	}

	return i, nil
}

// UpdateItemMenu atualiza os dados de um ItemMenu.
func (store *ItensMenuStore) UpdateItemMenu(item *models.ItemMenu) error {
	sqlString := fmt.Sprintf(updateItemMenuSQL, TableItensMenu, Nome, Descricao, Preco, ImagemURL, ItemID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdateItemMenu: %v", err)
		return fmt.Errorf("erro UpdateItemMenu: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Nome, item.Descricao, item.Preco, item.ImagemURL, item.ItemID)
	if err != nil {
		log.Printf("erro UpdateItemMenu: %v", err)
		return fmt.Errorf("erro UpdateItemMenu: %v", err)
	}

	return nil
}

// DeleteItemMenu remove um ItemMenu do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *ItensMenuStore) DeleteItemMenu(id int) error {
	sqlString := fmt.Sprintf(deleteItemMenuSQL, TableItensMenu, ItemID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro DeleteItemMenu: %v", err)
		return fmt.Errorf("erro DeleteItemMenu: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("erro DeleteItemMenu: %v", err)
		return fmt.Errorf("erro DeleteItemMenu: %v", err)
	}

	return nil
}

// GetAllItemMenu busca todos os itens do menu.
func (store *ItensMenuStore) GetAllItemMenu() ([]*models.ItemMenu, error) {
	sqlString := fmt.Sprintf(getALLItemMenuSQL, ItemID, Nome, Descricao, Preco, ImagemURL,
		TableItensMenu, Nome)

	rows, err := store.DB.Query(sqlString)
	if err != nil {
		log.Printf("erro GetAllItemMenu: %v", err)
		return nil, fmt.Errorf("erro GetAllItemMenu: %v", err)
	}
	defer rows.Close()

	var itensMenu []*models.ItemMenu
	for rows.Next() {
		var item models.ItemMenu
		if err := rows.Scan(&item.ItemID, &item.Nome, &item.Descricao, &item.Preco, &item.ImagemURL); err != nil {
			log.Printf("erro GetAllItemMenu: %v", err)
			return nil, fmt.Errorf("erro GetAllItemMenu: %v", err)
		}
		itensMenu = append(itensMenu, &item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetAllItemMenu: %v", err)
		return nil, fmt.Errorf("erro GetAllItemMenu: %v", err)
	}

	return itensMenu, nil
}

// GetItemMenuByNome busca por um ItemMenu pelo Nome
func (store *ItensMenuStore) GetItemMenuByNome(nome string) (*models.ItemMenu, error) {
	item := &models.ItemMenu{}

	sqlString := fmt.Sprintf(getItemMenuByNameSQL, ItemID, Nome, Descricao, Preco, ImagemURL,
		TableItensMenu, Nome)

	err := store.DB.QueryRow(sqlString, nome).Scan(
		&item.ItemID, &item.Nome, &item.Descricao, &item.Preco, &item.ImagemURL)
	if err != nil {
		log.Printf("erro GetItemMenuByNome: %v", err)
		return nil, fmt.Errorf("erro GetItemMenuByNome: %v", err)
	}

	return item, nil
}
