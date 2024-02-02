package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UsuarioHandler gerencia as requisições HTTP para usuários.
type UsuarioHandler struct {
	Service *services.UsuarioService
}

// NewUsuarioHandler cria uma nova instância de UsuarioHandler.
func NewUsuarioHandler(service *services.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{
		Service: service,
	}
}

// CreateUsuario lida com requisições POST para adicionar um novo usuário.
func (handler *UsuarioHandler) CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.CreateUsuario(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// GetUsuario lida com requisições GET para buscar um usuário pelo ID.
func (handler *UsuarioHandler) GetUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Service.GetUsuario(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.SenhaHash = "" // mantém informações confidenciais no servidor

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUsuario lida com requisições PUT para atualizar um usuário existente.
func (handler *UsuarioHandler) UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateUsuario(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

// DeleteUsuario lida com requisições DELETE para remover um usuário.
func (handler *UsuarioHandler) DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeleteUsuario(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
