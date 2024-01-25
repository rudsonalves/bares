package services

import (
	"bares_api/models"
	"bares_api/store"
)

// ItemPedidoService mantém a conexão com o banco de dados para operações
// relacionadas a itens pedidos.
type ItemPedidoService struct {
	store store.ItemPedidoStorer
}

// NewItemPedidoService cria uma nova instância de ItemPedidoService.
func NewItemPedidoService(store store.ItemPedidoStorer) *ItemPedidoService {
	return &ItemPedidoService{
		store: store,
	}
}

// CreateItemPedido adiciona um novo ItemPedido ao banco de dados.
func (service *ItemPedidoService) CreateItemPedido(item *models.ItemPedido) error {
	return service.store.CreateItemPedidoStore(item)
}

// GetItemPedido busca um itemPedido pelo ID.
func (service *ItemPedidoService) GetItemPedido(id int) (*models.ItemPedido, error) {
	return service.store.GetItemPedidoStore(id)
}

// UpdateItemPedido atualiza os dados de um itemPedido.
func (service *ItemPedidoService) UpdateItemPedido(item *models.ItemPedido) error {
	return service.store.UpdateItemPedidoStore(item)
}

// DeleteItemPedido remove um itemPedido do banco de dados.
func (service *ItemPedidoService) DeleteItemPedido(id int) error {
	return service.store.DeleteItemPedidoStore(id)
}
