package services

import (
	"bares_api/models"
	"bares_api/store"
)

// PedidoService fornece métodos para operações relacionadas a Pedido.
type PedidoService struct {
	store store.PedidoStorer
}

// NewPedidoService cria uma nova instância de PedidoService.
func NewPedidoService(store store.PedidoStorer) *PedidoService {
	return &PedidoService{
		store: store,
	}
}

// CreatePedido trata da lógica de negócios para criar um novo Pedido.
func (service *PedidoService) CreatePedido(pedido *models.Pedido) error {
	return service.store.CreatePedidoStore(pedido)
}

// GetPedido trata da lógica para recuperar um Pedido pelo ID.
func (service *PedidoService) GetPedido(id int) (*models.Pedido, error) {
	return service.store.GetPedidoStore(id)
}

// UpdatePedido atualiza os dados de um pedido.
func (service *PedidoService) UpdatePedido(pedido *models.Pedido) error {
	return service.store.UpdatePedidoStore(pedido)
}

// DeletePedido remove um pedido do banco de dados.
func (service *PedidoService) DeletePedido(id int) error {
	return service.store.DeletePedidoStore(id)
}

// GetPedidosByUsuario busca todos os pedidos de um usuário específico pelo usuarioID.
func (service *PedidoService) GetPedidosByUsuario(usuarioId int) ([]*models.Pedido, error) {
	return service.store.GetPedidosByUsuarioStore(usuarioId)
}
