package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MenuItemHandler manages HTTP requests for itemMenu.
type MenuItemHandler struct {
	Service *services.MenuItemService
}

// NewMenuItemHandler creates a new instance of MenuItemHandler.
func NewMenuItemHandler(service *services.MenuItemService) *MenuItemHandler {
	return &MenuItemHandler{
		Service: service,
	}
}

// CreateMenuItem handles POST requests to add a new ItemMenu.
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
	if err := json.NewEncoder(w).Encode(menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Get MenuItem handle GET requests to search for a Menu Item by ID.
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
	if err := json.NewEncoder(w).Encode(itemMenu); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Update MenuItem handle PUT requests to update an existing Menu Item.
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
	if err := json.NewEncoder(w).Encode(menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Delete MenuItem handle DELETE requests to remove a Menu Item.
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

// GetAllMenuItem handles GET requests to fetch all menu items.
func (handler *MenuItemHandler) GetAllMenuItem(w http.ResponseWriter, r *http.Request) {
	itensMenu, err := handler.Service.GetAllMenuItem()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(itensMenu); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetMenuItemByName handles GET requests to return an item by name
func (handler *MenuItemHandler) GetMenuItemByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nome, ok := vars["name"]
	if !ok {
		http.Error(w, "GetItemMenuByName error", http.StatusBadRequest)
		return
	}

	itemMenu, err := handler.Service.GetMenuItemByNome(nome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(itemMenu); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
