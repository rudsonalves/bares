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

// OrderStore mantém a conexão com o banco de dados para operações relacionadas a pedidos.
type OrderStore struct {
	DB *sql.DB
}

// NewOrder cria uma nova instância de PedidoStore.
func NewOrder(db *sql.DB) *OrderStore {
	return &OrderStore{DB: db}
}

// OrderStorer define as operações que um PedidoStore precisa implementar.
type OrderStorer interface {
	CreateOrder(pedido *models.Order) error
	GetOrder(id int) (*models.Order, error)
	UpdateOrder(pedido *models.Order) error
	DeleteOrder(id int) error
	GetOrderByUser(usuarioID int) ([]*models.Order, error)
	GetPendingOrders() ([]*models.Order, error)
}

// Garanta que PedidoStore implementa PedidoStorer.
var _ OrderStorer = &OrderStore{}

// CreateOrder adiciona um novo pedido ao banco de dados.
func (store *OrderStore) CreateOrder(pedido *models.Order) error {
	sqlString := fmt.Sprintf(createOrderSQL, TableOrders, UserID, DateTime, Status)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro CreatePedido: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(pedido.UsuarioID, pedido.DataHora, pedido.Status)
	if err != nil {
		log.Printf("erro CreatePedido: %v", err)
		return err
	}

	pedidoID, err := result.LastInsertId()
	if err != nil {
		log.Printf("erro CreatePedido: %v", err)
		return err
	}
	pedido.PedidoID = int(pedidoID)

	return nil
}

// GetOrder busca um pedido pelo ID.
func (store *OrderStore) GetOrder(id int) (*models.Order, error) {
	p := &models.Order{}

	sqlString := fmt.Sprintf(getOrderSQL, UserID, DateTime, Status, TableOrders, OrderID)

	err := store.DB.QueryRow(sqlString, id).Scan(&p.UsuarioID, &p.DataHora, &p.Status)
	if err != nil {
		log.Printf("erro GetPedido: %v", err)
		return nil, err
	}

	return p, nil
}

// UpdateOrder atualiza os dados de um pedido.
func (store *OrderStore) UpdateOrder(pedido *models.Order) error {
	sqlString := fmt.Sprintf(updateOrderSQL, TableOrders, UserID, DateTime, Status, OrderID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdatePedido: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pedido.UsuarioID, pedido.DataHora, pedido.Status, pedido.PedidoID)
	if err != nil {
		log.Printf("erro UpdatePedido: %v", err)
		return err
	}
	return nil
}

// DeleteOrder remove um pedido do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *OrderStore) DeleteOrder(id int) error {
	sqlString := fmt.Sprintf(deleteOrderSQL, TableOrders, OrderID)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro DeletePedido: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("erro DeletePedido: %v", err)
		return err
	}

	return nil
}

// GetOrderByUser busca todos os pedidos de um usuário específico pelo UsuarioID.
func (store *OrderStore) GetOrderByUser(usuarioID int) ([]*models.Order, error) {
	var pedidos []*models.Order

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
		OrderID, UserID, DateTime, Status, TableOrders, UserID)

	rows, err := store.DB.Query(queryString, usuarioID)
	if err != nil {
		log.Printf("erro GetPedidosByUsuario: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dataHoraStr string
		p := &models.Order{}
		err := rows.Scan(&p.PedidoID, &p.UsuarioID, &dataHoraStr, &p.Status)
		if err != nil {
			log.Printf("erro GetPedidosByUsuario: %v", err)
			return nil, err
		}
		p.DataHora, err = dateHourParse(dataHoraStr)
		if err != nil {
			log.Printf("erro GetPedidosByUsuario: %v", err)
			return nil, err
		}

		pedidos = append(pedidos, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetPedidosByUsuario: %v", err)
		return nil, err
	}

	return pedidos, nil
}

// GetPendingOrders retorna os pedidos com status diferente de 'entregue'
func (store *OrderStore) GetPendingOrders() ([]*models.Order, error) {
	var pedidos []*models.Order

	sqlString := fmt.Sprintf(getOrderPendindSQL, OrderID, UserID, DateTime, Status, TableOrders, Status)

	rows, err := store.DB.Query(sqlString, models.Entregue)
	if err != nil {
		log.Printf("erro GetPedidosPending: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dataHoraStr string
		p := &models.Order{}
		err := rows.Scan(&p.PedidoID, &p.UsuarioID, &dataHoraStr, &p.Status)
		if err != nil {
			log.Printf("erro GetPedidosPending: %v", err)
			return nil, err
		}
		p.DataHora, err = dateHourParse(dataHoraStr)
		if err != nil {
			log.Printf("erro GetPedidosPending: %v", err)
			return nil, err
		}

		pedidos = append(pedidos, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetPedidosPending: %v", err)
		return nil, err
	}

	return pedidos, nil
}

// Parse parses a formatted string and returns the time value it represents.
func dateHourParse(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
}
