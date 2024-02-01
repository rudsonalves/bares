package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ItemPedidoHandler gerencia as requisições HTTP para ItemPedido.
type ItemPedidoHandler struct {
  Service *services.ItemPedidoService
}

// NewItemPedidoHandler cria uma nova instância de ItemPedidoHandler.
func NewItemPedidoHandler(service *services.ItemPedidoService) *ItemPedidoHandler {
  return &ItemPedidoHandler{
    Service: service,
  }
}

// CreateItemPedido lida com requisições POST para adicionar um novo ItemPedido.
func (handler *ItemPedidoHandler) CreateItemPedido(w http.ResponseWriter, r *http.Request) {
  var itemPedido models.ItemPedido
  if err := json.NewDecoder(r.Body).Decode(&itemPedido); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if err := handler.Service.CreateItemPedido(&itemPedido); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(itemPedido)
}

// GetItemPedido lida com requisições GET para buscar um ItemPedido pelo ID.
func (handler *ItemPedidoHandler) GetItemPedido(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  itemPedido, err := handler.Service.GetItemPedido(id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(itemPedido)
}

// UpdateItemPedido lida com requisições PUT para atualizar um ItemPedido existente.
func (handler *ItemPedidoHandler) UpdateItemPedido(w http.ResponseWriter, r *http.Request) {
  var itemPedido models.ItemPedido
  if err := json.NewDecoder(r.Body).Decode(&itemPedido); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if err := handler.Service.UpdateItemPedido(&itemPedido); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(itemPedido)
}

// DeleteItemPedido lida com requisições DELETE para remover um ItemPedido.
func (handler *ItemPedidoHandler) DeleteItemPedido(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if err := handler.Service.DeleteItemPedido(id); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusOK)
}
