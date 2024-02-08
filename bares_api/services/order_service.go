package services

import (
	"bares_api/models"
	"bares_api/store"
)

// OrderService fornece métodos para operações relacionadas a Pedido.
type OrderService struct {
	store store.OrderStorer
}

// NewPedidoService cria uma nova instância de PedidoService.
func NewPedidoService(store store.OrderStorer) *OrderService {
	return &OrderService{
		store: store,
	}
}

// CreateOrder trata da lógica de negócios para criar um novo Pedido.
func (service *OrderService) CreateOrder(pedido *models.Order) error {
	return service.store.CreateOrder(pedido)
}

// GetOrder trata da lógica para recuperar um Pedido pelo ID.
func (service *OrderService) GetOrder(id int) (*models.Order, error) {
	return service.store.GetOrder(id)
}

// UpdateOrder atualiza os dados de um pedido.
func (service *OrderService) UpdateOrder(pedido *models.Order) error {
	return service.store.UpdateOrder(pedido)
}

// DeleteOrder remove um pedido do banco de dados.
func (service *OrderService) DeleteOrder(id int) error {
	return service.store.DeleteOrder(id)
}

// GetOrderByUser busca todos os pedidos de um usuário específico pelo usuarioID.
func (service *OrderService) GetOrderByUser(usuarioId int) ([]*models.Order, error) {
	return service.store.GetOrderByUser(usuarioId)
}

// GetPendingOrder busca todos os pedidos de um usuário específico pelo usuarioID.
func (service *OrderService) GetPendingOrder() ([]*models.Order, error) {
	return service.store.GetPendingOrders()
}
