package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
)

// ItensMenuStore mantém a conexão com o banco de dados para operações relacionadas a itens do menu.
type ItensMenuStore struct {
	DB *sql.DB
}

// NewItensMenuStore cria uma nova instância de ItensMenuStorer.
func NewItensMenuStore(db *sql.DB) *ItensMenuStore {
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
	createString := fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)",
		TableItensMenu, Nome, Descricao, Preco, ImagemURL)

	stmt, err := store.DB.Prepare(createString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.Nome, item.Descricao, item.Preco, item.ImagemURL)
	if err != nil {
		return err
	}

	itemID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	item.ItemID = int(itemID)

	return nil
}

// GetItemMenu busca um ItemMenu pelo ID.
func (store *ItensMenuStore) GetItemMenu(id int) (*models.ItemMenu, error) {
	i := &models.ItemMenu{}

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		ItemID, Nome, Descricao, Preco, ImagemURL, TableItensMenu, ItemID)

	err := store.DB.QueryRow(queryString, id).Scan(&i.ItemID, &i.Nome, &i.Descricao, &i.Preco, &i.ImagemURL)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// UpdateItemMenu atualiza os dados de um ItemMenu.
func (store *ItensMenuStore) UpdateItemMenu(item *models.ItemMenu) error {
	updateString := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?",
		TableItensMenu, Nome, Descricao, Preco, ImagemURL, ItemID)

	stmt, err := store.DB.Prepare(updateString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Nome, item.Descricao, item.Preco, item.ImagemURL)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItemMenu remove um ItemMenu do banco de dados.
func (store *ItensMenuStore) DeleteItemMenu(id int) error {
	deleteString := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", TableItensMenu, ItemID)

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

// GetAllItemMenu busca todos os itens do menu.
func (store *ItensMenuStore) GetAllItemMenu() ([]*models.ItemMenu, error) {
	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s ORDER BY %s",
		ItemID, Nome, Descricao, Preco, ImagemURL, TableItensMenu, Nome)

	rows, err := store.DB.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itensMenu []*models.ItemMenu
	for rows.Next() {
		var item models.ItemMenu
		if err := rows.Scan(&item.ItemID, &item.Nome, &item.Descricao,
			&item.Preco, &item.ImagemURL); err != nil {
			return nil, err
		}
		itensMenu = append(itensMenu, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return itensMenu, nil
}

// GetItemMenuByNome busca por um ItemMenu pelo Nome
func (store *ItensMenuStore) GetItemMenuByNome(nome string) (*models.ItemMenu, error) {
	item := &models.ItemMenu{}

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		ItemID, Nome, Descricao, Preco, ImagemURL, TableItensMenu, Nome)

	err := store.DB.QueryRow(queryString, nome).Scan(
		&item.ItemID, &item.Nome, &item.Descricao, &item.Preco, &item.ImagemURL)
	if err != nil {
		return nil, err
	}

	return item, nil
}
