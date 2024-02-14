package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ItemOrderHandler gerencia as requisições HTTP para ItemPedido.
type ItemOrderHandler struct {
	Service *services.ItemOrderService
}

// NewItemOrderHandler cria uma nova instância de ItemPedidoHandler.
func NewItemOrderHandler(service *services.ItemOrderService) *ItemOrderHandler {
	return &ItemOrderHandler{
		Service: service,
	}
}

// CreateItemOrder lida com requisições POST para adicionar um novo ItemPedido.
func (handler *ItemOrderHandler) CreateItemOrder(w http.ResponseWriter, r *http.Request) {
	var itemOrder models.ItemOrder
	if err := json.NewDecoder(r.Body).Decode(&itemOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.CreateItemOrder(&itemOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(itemOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetIItemOrder lida com requisições GET para buscar um ItemPedido pelo ID.
func (handler *ItemOrderHandler) GetIItemOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemPedido, err := handler.Service.GetItemOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(itemPedido); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateItemOrder lida com requisições PUT para atualizar um ItemPedido existente.
func (handler *ItemOrderHandler) UpdateItemOrder(w http.ResponseWriter, r *http.Request) {
	var itemOrder models.ItemOrder
	if err := json.NewDecoder(r.Body).Decode(&itemOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateItemOrder(&itemOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(itemOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteItemOrder lida com requisições DELETE para remover um ItemPedido.
func (handler *ItemOrderHandler) DeleteItemOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeleteItemOrder(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
