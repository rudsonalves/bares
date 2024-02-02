package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	createOrderSQL     = "INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)"
	getOrderSQL        = "SELECT %s, %s, %s FROM %s WHERE %s = ?"
	getOrderPendindSQL = "SELECT %s, %s, %s, %s FROM %s WHERE %s != ?"
	updateOrderSQL     = "UPDATE %s SET %s = ?, %s = ?, %s = ? WHERE %s = ?"
	deleteOrderSQL     = "DELETE FROM %s WHERE %s = ?"
)

// PedidoStore mantém a conexão com o banco de dados para operações relacionadas a pedidos.
type PedidoStore struct {
	DB *sql.DB
}

// NewPedido cria uma nova instância de PedidoStore.
func NewPedido(db *sql.DB) *PedidoStore {
	return &PedidoStore{DB: db}
}

// PedidoStorer define as operações que um PedidoStore precisa implementar.
type PedidoStorer interface {
	CreatePedido(pedido *models.Pedido) error
	GetPedido(id int) (*models.Pedido, error)
	UpdatePedido(pedido *models.Pedido) error
	DeletePedido(id int) error
	GetPedidosByUsuario(usuarioID int) ([]*models.Pedido, error)
	GetPedidosPending() ([]*models.Pedido, error)
}

// Garanta que PedidoStore implementa PedidoStorer.
var _ PedidoStorer = &PedidoStore{}

// CreatePedido adiciona um novo pedido ao banco de dados.
func (store *PedidoStore) CreatePedido(pedido *models.Pedido) error {
	sqlString := fmt.Sprintf(createOrderSQL, TablePedidos, UsuarioID, DataHora, Status)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro CreatePedido: %v", err)
		return fmt.Errorf("erro CreatePedido: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(pedido.UsuarioID, pedido.DataHora, pedido.Status)
	if err != nil {
		log.Printf("erro CreatePedido: %v", err)
		return fmt.Errorf("erro CreatePedido: %v", err)
	}

	pedidoID, err := result.LastInsertId()
	if err != nil {
		log.Printf("erro CreatePedido: %v", err)
		return fmt.Errorf("erro CreatePedido: %v", err)
	}
	pedido.PedidoID = int(pedidoID)

	return nil
}

// GetPedido busca um pedido pelo ID.
func (store *PedidoStore) GetPedido(id int) (*models.Pedido, error) {
	p := &models.Pedido{}

	sqlString := fmt.Sprintf(getOrderSQL, UsuarioID, DataHora, Status, TablePedidos, PedidoID)

	err := store.DB.QueryRow(sqlString, id).Scan(&p.UsuarioID, &p.DataHora, &p.Status)
	if err != nil {
		log.Printf("erro GetPedido: %v", err)
		return nil, fmt.Errorf("erro GetPedido: %v", err)
	}

	return p, nil
}

// UpdatePedido atualiza os dados de um pedido.
func (store *PedidoStore) UpdatePedido(pedido *models.Pedido) error {
	sqlString := fmt.Sprintf(updateOrderSQL, TablePedidos, UsuarioID, DataHora, Status, PedidoID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdatePedido: %v", err)
		return fmt.Errorf("erro UpdatePedido: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(pedido.UsuarioID, pedido.DataHora, pedido.Status, pedido.PedidoID)
	if err != nil {
		log.Printf("erro UpdatePedido: %v", err)
		return fmt.Errorf("erro UpdatePedido: %v", err)
	}
	return nil
}

// DeletePedido remove um pedido do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *PedidoStore) DeletePedido(id int) error {
	sqlString := fmt.Sprintf(deleteOrderSQL, TablePedidos, PedidoID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro DeletePedido: %v", err)
		return fmt.Errorf("erro DeletePedido: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("erro DeletePedido: %v", err)
		return fmt.Errorf("erro DeletePedido: %v", err)
	}

	return nil
}

// GetPedidosByUsuario busca todos os pedidos de um usuário específico pelo UsuarioID.
func (store *PedidoStore) GetPedidosByUsuario(usuarioID int) ([]*models.Pedido, error) {
	var pedidos []*models.Pedido

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
		PedidoID, UsuarioID, DataHora, Status, TablePedidos, UsuarioID)

	rows, err := store.DB.Query(queryString, usuarioID)
	if err != nil {
		log.Printf("erro GetPedidosByUsuario: %v", err)
		return nil, fmt.Errorf("erro GetPedidosByUsuario: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataHoraStr string
		p := &models.Pedido{}
		err := rows.Scan(&p.PedidoID, &p.UsuarioID, &dataHoraStr, &p.Status)
		if err != nil {
			log.Printf("erro GetPedidosByUsuario: %v", err)
			return nil, fmt.Errorf("erro GetPedidosByUsuario: %v", err)
		}
		p.DataHora, err = dateHourParse(dataHoraStr)
		if err != nil {
			log.Printf("erro GetPedidosByUsuario: %v", err)
			return nil, fmt.Errorf("erro GetPedidosByUsuario: %v", err)
		}

		pedidos = append(pedidos, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetPedidosByUsuario: %v", err)
		return nil, fmt.Errorf("erro GetPedidosByUsuario: %v", err)
	}

	return pedidos, nil
}

// GetPedidosPending retorna os pedidos com status diferente de 'entregue'
func (store *PedidoStore) GetPedidosPending() ([]*models.Pedido, error) {
	var pedidos []*models.Pedido

	sqlString := fmt.Sprintf(getOrderPendindSQL, PedidoID, UsuarioID, DataHora, Status, TablePedidos, Status)

	rows, err := store.DB.Query(sqlString, models.Entregue)
	if err != nil {
		log.Printf("erro GetPedidosPending: %v", err)
		return nil, fmt.Errorf("erro GetPedidosPending: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataHoraStr string
		p := &models.Pedido{}
		err := rows.Scan(&p.PedidoID, &p.UsuarioID, &dataHoraStr, &p.Status)
		if err != nil {
			log.Printf("erro GetPedidosPending: %v", err)
			return nil, fmt.Errorf("erro GetPedidosPending: %v", err)
		}
		p.DataHora, err = dateHourParse(dataHoraStr)
		if err != nil {
			log.Printf("erro GetPedidosPending: %v", err)
			return nil, fmt.Errorf("erro GetPedidosPending: %v", err)
		}

		pedidos = append(pedidos, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetPedidosPending: %v", err)
		return nil, fmt.Errorf("erro GetPedidosPending: %v", err)
	}

	return pedidos, nil
}

// Parse parses a formatted string and returns the time value it represents.
func dateHourParse(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
}
