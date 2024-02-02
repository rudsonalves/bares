package services

import (
	"bares_api/models"
	"bares_api/store"
	"fmt"
)

// ItemMenuService fornece métodos para operações relacionadas a ItemMenu.
type ItemMenuService struct {
	store store.ItensMenuStorer
}

// NewItemMenuService cria uma nova instância de ItemMenuService.
func NewItemMenuService(store store.ItensMenuStorer) *ItemMenuService {
	return &ItemMenuService{
		store: store,
	}
}

// CreateItemMenu trata da lógica de negócios para criar um novo ItemMenu.
func (service *ItemMenuService) CreateItemMenu(item *models.ItemMenu) error {
	// Verificar se o Nome já está em uso:
	existingItem, err := service.store.GetItemMenuByNome(item.Nome)
	if err != nil {
		return err
	}
	if existingItem != nil {
		return fmt.Errorf("item '%s' já existe em ItensMenu", item.Nome)
	}

	// Continuar com a criação do ItemMenu
	return service.store.CreateItemMenu(item)
}

// GetItemMenu para buscar um ItemMenu pelo ID.
func (service *ItemMenuService) GetItemMenu(id int) (*models.ItemMenu, error) {
	return service.store.GetItemMenu(id)
}

// UpdateItemMenu trata da lógica para atualizar um ItemMenu existente.
func (service *ItemMenuService) UpdateItemMenu(item *models.ItemMenu) error {
	return service.store.UpdateItemMenu(item)
}

// DeleteItemMenu trata da lógica para deletar um ItemMenu.
func (service *ItemMenuService) DeleteItemMenu(id int) error {
	return service.store.DeleteItemMenu(id)
}

// GetAllItemMenu busca todos os itens do menu.
func (service *ItemMenuService) GetAllItemMenu() ([]*models.ItemMenu, error) {
	return service.store.GetAllItemMenu()
}

// GetItemMenuByNome retorna um item pelo nome
func (service *ItemMenuService) GetItemMenuByNome(nome string) (*models.ItemMenu, error) {
	return service.store.GetItemMenuByNome(nome)
}
