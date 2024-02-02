package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PedidoHandler gerencia as requisições HTTP para Pedido.
type PedidoHandler struct {
	Service *services.PedidoService
}

// NewPedidoHandler cria uma nova instância de PedidoHandler.
func NewPedidoHandler(service *services.PedidoService) *PedidoHandler {
	return &PedidoHandler{
		Service: service,
	}
}

// CreatePedido lida com requisições POST para adicionar um novo Pedido.
func (handler *PedidoHandler) CreatePedido(w http.ResponseWriter, r *http.Request) {
	var pedido models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.CreatePedido(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pedido)
}

// GetPedido lida com requisições GET para buscar um Pedido pelo ID.
func (handler *PedidoHandler) GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido, err := handler.Service.GetPedido(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

// UpdatePedido lida com requisições PUT para atualizar um Pedido existente.
func (handler *PedidoHandler) UpdatePedido(w http.ResponseWriter, r *http.Request) {
	var pedido models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdatePedido(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedido)
}

// DeletePedido lida com requisições DELETE para remover um Pedido.
func (handler *PedidoHandler) DeletePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeletePedido(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetPedidosByUsuario lida com requisições GET para buscar um Pedido pelo usuarioId.
func (handler *PedidoHandler) GetPedidosByUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedidos, err := handler.Service.GetPedidosByUsuario(usuarioId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedidos)
}

// GetPedidosPending busca todos os pedidos de um usuário específico pelo usuarioID.
func (handler *PedidoHandler) GetPedidosPending(w http.ResponseWriter, r *http.Request) {
	pedidos, err := handler.Service.GetPedidosPending()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedidos)
}
