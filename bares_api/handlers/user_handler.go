package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserHandler manages HTTP requests for users.
type UserHandler struct {
	Service *services.UserService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// CreateUser handles POST requests to add a new user.
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRoleFromToken, _ := r.Context().Value("Role").(models.Role)
	if userRoleFromToken == models.Garcom && user.Role != models.Cliente {
		http.Error(w, "Garçons só podem criar usuários do tipo Cliente.", http.StatusInternalServerError)
		return
	}

	if err := handler.Service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser handles GET requests to look up a user by ID.
func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.PasswordHash = "" // keeps confidential information on the server

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser handles PUT requests to update an existing user.
func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var usuario models.User
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateUser(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

// DeleteUser handles DELETE requests to remove a user.
func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
