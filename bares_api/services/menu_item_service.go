package services

import (
	"bares_api/models"
	"bares_api/store"
	"fmt"
)

// MenuItemService fornece métodos para operações relacionadas a ItemMenu.
type MenuItemService struct {
	store store.MenuItemStorer
}

// NewItemMenuService cria uma nova instância de ItemMenuService.
func NewItemMenuService(store store.MenuItemStorer) *MenuItemService {
	return &MenuItemService{
		store: store,
	}
}

// CreateMenuItem trata da lógica de negócios para criar um novo ItemMenu.
func (service *MenuItemService) CreateMenuItem(item *models.MenuItem) error {
	// Verificar se o Nome já está em uso:
	existingItem, err := service.store.GetMenuItemByName(item.Nome)
	if err != nil {
		return err
	}
	if existingItem != nil {
		return fmt.Errorf("item '%s' já existe em ItensMenu", item.Nome)
	}

	// Continuar com a criação do ItemMenu
	return service.store.CreateMenuItem(item)
}

// GetMenuItem para buscar um ItemMenu pelo ID.
func (service *MenuItemService) GetMenuItem(id int) (*models.MenuItem, error) {
	return service.store.GetMenuItem(id)
}

// UpdateMenuItem trata da lógica para atualizar um ItemMenu existente.
func (service *MenuItemService) UpdateMenuItem(item *models.MenuItem) error {
	return service.store.UpdateMenuItem(item)
}

// DeleteMenuItem trata da lógica para deletar um ItemMenu.
func (service *MenuItemService) DeleteMenuItem(id int) error {
	return service.store.DeleteMenuItem(id)
}

// GetAllMenuItem busca todos os itens do menu.
func (service *MenuItemService) GetAllMenuItem() ([]*models.MenuItem, error) {
	return service.store.GetAllMenuItem()
}

// GetMenuItemByNome retorna um item pelo nome
func (service *MenuItemService) GetMenuItemByNome(nome string) (*models.MenuItem, error) {
	return service.store.GetMenuItemByName(nome)
}
