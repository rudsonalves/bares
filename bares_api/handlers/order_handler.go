package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// OrderHandler gerencia as requisições HTTP para Pedido.
type OrderHandler struct {
	Service *services.OrderService
}

// NewOrderHandler cria uma nova instância de PedidoHandler.
func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{
		Service: service,
	}
}

// CreateOrder lida com requisições POST para adicionar um novo Pedido.
func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var pedido models.Order
	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.CreateOrder(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pedido)
}

// GetOrder lida com requisições GET para buscar um Pedido pelo ID.
func (handler *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido, err := handler.Service.GetOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

// UpdateOrder lida com requisições PUT para atualizar um Pedido existente.
func (handler *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// DeleteOrder lida com requisições DELETE para remover um Pedido.
func (handler *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeleteOrder(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetOrderByUser lida com requisições GET para buscar um Pedido pelo usuarioId.
func (handler *OrderHandler) GetOrderByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedidos, err := handler.Service.GetOrderByUser(usuarioId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedidos)
}

// GetPendingOrder busca todos os pedidos de um usuário específico pelo usuarioID.
func (handler *OrderHandler) GetPendingOrder(w http.ResponseWriter, r *http.Request) {
	pedidos, err := handler.Service.GetPendingOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedidos)
}
