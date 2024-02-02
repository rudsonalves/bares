package services

import (
	"bares_api/models"
	"bares_api/store"
)

// ItemOrderService mantém a conexão com o banco de dados para operações
// relacionadas a itens pedidos.
type ItemOrderService struct {
	store store.ItemOrderStorer
}

// NewItemPedidoService cria uma nova instância de ItemPedidoService.
func NewItemPedidoService(store store.ItemOrderStorer) *ItemOrderService {
	return &ItemOrderService{
		store: store,
	}
}

// CreateItemOrder adiciona um novo ItemPedido ao banco de dados.
func (service *ItemOrderService) CreateItemOrder(item *models.ItemOrder) error {
	return service.store.CreateItemOrder(item)
}

// GetItemOrder busca um itemPedido pelo ID.
func (service *ItemOrderService) GetItemOrder(id int) (*models.ItemOrder, error) {
	return service.store.GetItemOrder(id)
}

// UpdateItemOrder atualiza os dados de um itemPedido.
func (service *ItemOrderService) UpdateItemOrder(item *models.ItemOrder) error {
	return service.store.UpdateItemOrder(item)
}

// DeleteItemOrder remove um itemPedido do banco de dados.
func (service *ItemOrderService) DeleteItemOrder(id int) error {
	return service.store.DeleteItemOrder(id)
}
