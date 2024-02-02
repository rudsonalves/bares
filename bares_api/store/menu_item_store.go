package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"
)

const (
	createItemMenuSQL    = "INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)"
	getItemMenuSQL       = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
	updateItemMenuSQL    = "UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?"
	deleteItemMenuSQL    = "DELETE FROM %s WHERE %s = ?"
	getALLItemMenuSQL    = "SELECT %s, %s, %s, %s, %s FROM %s ORDER BY %s"
	getItemMenuByNameSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
)

// MenuItemStore mantém a conexão com o banco de dados para operações relacionadas a itens do menu.
type MenuItemStore struct {
	DB *sql.DB
}

// NewMenuItem cria uma nova instância de MenuItemStorer.
func NewMenuItem(db *sql.DB) *MenuItemStore {
	return &MenuItemStore{DB: db}
}

// MenuItemStorer define as operações que um MenuItemStore precisa implementar.
type MenuItemStorer interface {
	CreateMenuItem(item *models.MenuItem) error
	GetMenuItem(id int) (*models.MenuItem, error)
	UpdateMenuItem(item *models.MenuItem) error
	DeleteMenuItem(id int) error
	GetAllMenuItem() ([]*models.MenuItem, error)
	GetMenuItemByName(nome string) (*models.MenuItem, error)
}

// Garanta que ItensMenuStore implementa MenuItemStorer.
var _ MenuItemStorer = &MenuItemStore{}

// CreateMenuItem adiciona um novo usuário ao banco de dados.
func (store *MenuItemStore) CreateMenuItem(item *models.MenuItem) error {
	sqlString := fmt.Sprintf(createItemMenuSQL,
		TableMenuItem, Name, Description, Price, ImagemURL)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro CreateItemMenu: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.Nome, item.Descricao, item.Preco, item.ImagemURL)
	if err != nil {
		log.Printf("erro CreateItemMenu: %v", err)
		return err
	}

	itemID, err := result.LastInsertId()
	if err != nil {
		log.Printf("erro CreateItemMenu: %v", err)
		return err
	}
	item.ItemID = int(itemID)

	return nil
}

// GetMenuItem busca um ItemMenu pelo ID.
func (store *MenuItemStore) GetMenuItem(id int) (*models.MenuItem, error) {
	i := &models.MenuItem{}

	sqlString := fmt.Sprintf(getItemMenuSQL, ItemID, Name, Description, Price, ImagemURL,
		TableMenuItem, ItemID)

	err := store.DB.QueryRow(sqlString, id).Scan(&i.ItemID, &i.Nome, &i.Descricao, &i.Preco, &i.ImagemURL)
	if err != nil {
		log.Printf("erro GetItemMenu: %v", err)
		return nil, err
	}

	return i, nil
}

// UpdateMenuItem atualiza os dados de um ItemMenu.
func (store *MenuItemStore) UpdateMenuItem(item *models.MenuItem) error {
	sqlString := fmt.Sprintf(updateItemMenuSQL, TableMenuItem, Name, Description, Price, ImagemURL, ItemID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdateItemMenu: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Nome, item.Descricao, item.Preco, item.ImagemURL, item.ItemID)
	if err != nil {
		log.Printf("erro UpdateItemMenu: %v", err)
		return err
	}

	return nil
}

// DeleteMenuItem remove um ItemMenu do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *MenuItemStore) DeleteMenuItem(id int) error {
	sqlString := fmt.Sprintf(deleteItemMenuSQL, TableMenuItem, ItemID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro DeleteItemMenu: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("erro DeleteItemMenu: %v", err)
		return err
	}

	return nil
}

// GetAllMenuItem busca todos os itens do menu.
func (store *MenuItemStore) GetAllMenuItem() ([]*models.MenuItem, error) {
	sqlString := fmt.Sprintf(getALLItemMenuSQL, ItemID, Name, Description, Price, ImagemURL,
		TableMenuItem, Name)

	rows, err := store.DB.Query(sqlString)
	if err != nil {
		log.Printf("erro GetAllItemMenu: %v", err)
		return nil, err
	}
	defer rows.Close()

	var itensMenu []*models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		if err := rows.Scan(&item.ItemID, &item.Nome, &item.Descricao, &item.Preco, &item.ImagemURL); err != nil {
			log.Printf("erro GetAllItemMenu: %v", err)
			return nil, err
		}
		itensMenu = append(itensMenu, &item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetAllItemMenu: %v", err)
		return nil, err
	}

	return itensMenu, nil
}

// GetMenuItemByName busca por um ItemMenu pelo Nome
func (store *MenuItemStore) GetMenuItemByName(nome string) (*models.MenuItem, error) {
	item := &models.MenuItem{}

	sqlString := fmt.Sprintf(getItemMenuByNameSQL, ItemID, Name, Description, Price, ImagemURL,
		TableMenuItem, Name)

	err := store.DB.QueryRow(sqlString, nome).Scan(
		&item.ItemID, &item.Nome, &item.Descricao, &item.Preco, &item.ImagemURL)
	if err != nil {
		log.Printf("erro GetItemMenuByNome: %v", err)
		return nil, err
	}

	return item, nil
}
