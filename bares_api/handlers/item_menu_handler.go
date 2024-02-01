package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ItemMenuHandler gerencia as requisições HTTP para itemMenu.
type ItemMenuHandler struct {
  Service *services.ItemMenuService
}

// NewItemMenuHandler cria uma nova instância de ItemMenuHandler.
func NewItemMenuHandler(service *services.ItemMenuService) *ItemMenuHandler {
  return &ItemMenuHandler{
    Service: service,
  }
}

// CreateItemMenu lida com requisições POST para adicionar um novo ItemMenu.
func (handler *ItemMenuHandler) CreateItemMenu(w http.ResponseWriter, r *http.Request) {
  var itemMenu models.ItemMenu
  if err := json.NewDecoder(r.Body).Decode(&itemMenu); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if err := handler.Service.CreateItemMenu(&itemMenu); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(itemMenu)
}

// GetItemMenu lida com requisições GET para buscar um ItemMenu pelo ID.
func (handler *ItemMenuHandler) GetItemMenu(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  itemMenu, err := handler.Service.GetItemMenu(id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(itemMenu)
}

// UpdateItemMenu lida com requisições PUT para atualizar um ItemMenu existente.
func (handler *ItemMenuHandler) UpdateItemMenu(w http.ResponseWriter, r *http.Request) {
  var itemMenu models.ItemMenu
  if err := json.NewDecoder(r.Body).Decode(&itemMenu); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if err := handler.Service.UpdateItemMenu(&itemMenu); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(itemMenu)
}

// DeleteItemMenu lida com requisições DELETE para remover um ItemMenu.
func (handler *ItemMenuHandler) DeleteItemMenu(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if err := handler.Service.DeleteItemMenu(id); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
}

// GetAllItemMenu lida com requisições GET para busca todos os itens do menu.
func (handler *ItemMenuHandler) GetAllItemMenu(w http.ResponseWriter, r *http.Request) {
  itensMenu, err := handler.Service.GetAllItemMenu()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(itensMenu)
}

// GetItemMenuByNome lida com requisições GET para retorna um item pelo nome
func (handler *ItemMenuHandler) GetItemMenuByNome(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  nome, ok := vars["name"]
  if !ok {
    http.Error(w, "GetItemMenuByNome error", http.StatusBadRequest)
    return
  }

  itemMenu, err := handler.Service.GetItemMenuByNome(nome)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(itemMenu)
}
