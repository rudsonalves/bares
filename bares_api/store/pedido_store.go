package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
)

// PedidoStore mantém a conexão com o banco de dados para operações relacionadas a pedidos.
type PedidoStore struct {
	DB *sql.DB
}

// NewPedidoStore cria uma nova instância de PedidoStore.
func NewPedidoStore(db *sql.DB) *PedidoStore {
	return &PedidoStore{DB: db}
}

// PedidoStorer define as operações que um PedidoStore precisa implementar.
type PedidoStorer interface {
	CreatePedidoStore(pedido *models.Pedido) error
	GetPedidoStore(id int) (*models.Pedido, error)
	UpdatePedidoStore(pedido *models.Pedido) error
	DeletePedidoStore(id int) error
	GetPedidosByUsuarioStore(usuarioID int) ([]*models.Pedido, error)
}

// Garanta que PedidoStore implementa PedidoStorer.
var _ PedidoStorer = &PedidoStore{}

// CreatePedidoStore adiciona um novo pedido ao banco de dados.
func (store *PedidoStore) CreatePedidoStore(pedido *models.Pedido) error {
	createString := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)",
		TablePedidos, UsuarioID, DataHora, Status)

	stmt, err := store.DB.Prepare(createString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(pedido.UsuarioID, pedido.DataHora, pedido.Status)
	if err != nil {
		return err
	}

	pedidoID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	pedido.PedidoID = int(pedidoID)

	return nil
}

// GetPedidoStore busca um pedido pelo ID.
func (store *PedidoStore) GetPedidoStore(id int) (*models.Pedido, error) {
	p := &models.Pedido{}

	queryString := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = ?",
		UsuarioID, DataHora, Status, TablePedidos, PedidoID)

	err := store.DB.QueryRow(queryString, id).Scan(&p.UsuarioID, &p.DataHora, &p.Status)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// UpdatePedidoStore atualiza os dados de um pedido.
func (store *PedidoStore) UpdatePedidoStore(pedido *models.Pedido) error {
	updateString := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ? WHERE %s = ?",
		TablePedidos, UsuarioID, DataHora, Status, PedidoID)

	stmt, err := store.DB.Prepare(updateString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pedido.UsuarioID, pedido.DataHora, pedido.Status, pedido.PedidoID)
	if err != nil {
		return err
	}
	return nil
}

// DeletePedidoStore remove um pedido do banco de dados.
func (store *PedidoStore) DeletePedidoStore(id int) error {
	deleteString := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", TablePedidos, PedidoID)

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

// GetPedidosByUsuarioStore busca todos os pedidos de um usuário específico pelo UsuarioID.
func (store *PedidoStore) GetPedidosByUsuarioStore(usuarioID int) ([]*models.Pedido, error) {
	var pedidos []*models.Pedido

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
		PedidoID, UsuarioID, DataHora, Status, TablePedidos, UsuarioID)

	rows, err := store.DB.Query(queryString, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &models.Pedido{}
		err := rows.Scan(&p.PedidoID, &p.UsuarioID, &p.DataHora, &p.Status)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pedidos, nil
}
