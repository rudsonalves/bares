package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MenuItemHandler gerencia as requisições HTTP para itemMenu.
type MenuItemHandler struct {
	Service *services.MenuItemService
}

// NewMenuItemHandler cria uma nova instância de MenuItemHandler.
func NewMenuItemHandler(service *services.MenuItemService) *MenuItemHandler {
	return &MenuItemHandler{
		Service: service,
	}
}

// CreateMenuItem lida com requisições POST para adicionar um novo ItemMenu.
func (handler *MenuItemHandler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	var menuItem models.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.CreateMenuItem(&menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menuItem)
}

// GetMenuItem lida com requisições GET para buscar um ItemMenu pelo ID.
func (handler *MenuItemHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemMenu, err := handler.Service.GetMenuItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemMenu)
}

// UpdateMenuItem lida com requisições PUT para atualizar um ItemMenu existente.
func (handler *MenuItemHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	var menuItem models.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateMenuItem(&menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuItem)
}

// DeleteMenuItem lida com requisições DELETE para remover um ItemMenu.
func (handler *MenuItemHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeleteMenuItem(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetAllMenuItem lida com requisições GET para busca todos os itens do menu.
func (handler *MenuItemHandler) GetAllMenuItem(w http.ResponseWriter, r *http.Request) {
	itensMenu, err := handler.Service.GetAllMenuItem()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itensMenu)
}

// GetMenuItemByNome lida com requisições GET para retorna um item pelo nome
func (handler *MenuItemHandler) GetMenuItemByNome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nome, ok := vars["name"]
	if !ok {
		http.Error(w, "GetItemMenuByNome error", http.StatusBadRequest)
		return
	}

	itemMenu, err := handler.Service.GetMenuItemByNome(nome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemMenu)
}
