package handlers

import (
	"bares_api/models"
	"bares_api/services"
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("UserHandler.CreateUser: starting...")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("UserHandler.GetAllUsers 0: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// userRoleFromToken, _ := r.Context().Value("Role").(models.Role)
	// if userRoleFromToken == models.Garcom && user.Role != models.Cliente {
	// 	http.Error(w, "Garçons só podem criar usuários do tipo Cliente.", http.StatusInternalServerError)
	// 	return
	// }

	if err := handler.Service.CreateUser(&user); err != nil {
		log.Printf("UserHandler.GetAllUsers 0: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("UserHandler.GetAllUsers 1: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetUser handles GET requests to look up a user by ID.
func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("UserHandler.GetUser: starting...")
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("UserHandler.GetUser 0: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Service.GetUser(id)
	if err != nil {
		fmt.Printf("UserHandler.GetUser 1: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.PasswordHash = "" // keeps confidential information on the server

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Printf("UserHandler.GetUser 2: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandlerGetAllUsers: starting...")
	users, err := handler.Service.GetAllUsers()
	if err != nil {
		log.Printf("UserHandler.GetAllUsers 0: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for index := range users {
		users[index].PasswordHash = "" // keeps confidential information on the server
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("UserHandler.GetAllUsers 1: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateUser handles PUT requests to update an existing user.
func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandler.UpdateUser: starting...")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("UserHandler.UpdateUser 0: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateUser(&user); err != nil {
		log.Printf("UserHandler.UpdateUser 1: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("UserHandler.UpdateUser 2: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *UserHandler) UpdateUserPass(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandler.UpdateUserPass 0: starting...")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("UserHandler.UpdateUserPass 1: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateUserPass(user.Id, user.PasswordHash); err != nil {
		log.Printf("UserHandler.UpdateUserPass 2: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser handles DELETE requests to remove a user.
func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandler.DeleteUser: starting...")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("UserHandler.DeleteUser 0: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.Service.DeleteUser(id); err != nil {
		log.Printf("UserHandler.DeleteUser 1: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
