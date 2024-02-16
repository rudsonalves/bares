package main

import (
	"bares_api/bootstrap"
	"bares_api/handlers"
	"bares_api/services"
	"bares_api/store"
	"bares_api/utils"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	const dbName = "BarDB" // Database name

	// Get username and password and connect to the database
	dbStore := utils.CreateDataBaseConn(dbName)
	defer dbStore.DatabaseClose()

	err := dbStore.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Starting stores
	userStore := store.NewUser(dbStore.DB)
	menuItemStore := store.NewMenuItem(dbStore.DB)
	orderStore := store.NewOrder(dbStore.DB)
	itemOrderStore := store.NewItemOrder(dbStore.DB)

	// Starting services
	userService := services.NewUsuarioService(userStore)
	authService := services.NewAuthservice(userStore)
	menuItemService := services.NewItemMenuService(menuItemStore)
	orderService := services.NewPedidoService(orderStore)
	itemOrderService := services.NewItemPedidoService(itemOrderStore)

	// Starting handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	menuItemHandler := handlers.NewMenuItemHandler(menuItemService)
	orderHandler := handlers.NewOrderHandler(orderService)
	itemOrderHandler := handlers.NewItemOrderHandler(itemOrderService)

	// Check if there is an 'admin' user or 'gerente' user.
	if err := bootstrap.CheckAndCreateAdminUser(userService); err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	// Configure routes
	route := mux.NewRouter()

	// Public Routes
	route.HandleFunc("/login", authHandler.LoginHandlers).Methods("POST")

	// Private Routes
	api := route.PathPrefix("/api").Subrouter()
	api.Use(handlers.AuthMiddleware) // Apply authentication middleware

	// The routes below require authentication
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/menuitem", menuItemHandler.CreateMenuItem).Methods("POST")
	api.HandleFunc("/order", orderHandler.CreateOrder).Methods("POST")
	api.HandleFunc("/itemorder", itemOrderHandler.CreateItemOrder).Methods("POST")

	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	api.HandleFunc("/menuitem", menuItemHandler.GetAllMenuItem).Methods("GET")
	api.HandleFunc("/menuitem/{id}", menuItemHandler.GetMenuItem).Methods("GET")
	api.HandleFunc("/menuitem/name/{name}", menuItemHandler.GetMenuItemByName).Methods("GET")
	api.HandleFunc("/order", orderHandler.GetPendingOrder).Methods("GET")
	api.HandleFunc("/order/{id}", orderHandler.GetOrder).Methods("GET")
	api.HandleFunc("/order/user/{id}", orderHandler.GetOrderByUser).Methods("GET")
	api.HandleFunc("/itemorder/{id}", itemOrderHandler.GetIItemOrder).Methods("GET")

	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/password/{id}", userHandler.UpdateUserPass).Methods("PUT")
	api.HandleFunc("/menuitem/{id}", menuItemHandler.UpdateMenuItem).Methods("PUT")
	api.HandleFunc("/order/{id}", orderHandler.UpdateOrder).Methods("PUT")
	api.HandleFunc("/itemorder/{id}", itemOrderHandler.UpdateItemOrder).Methods("PUT")

	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	api.HandleFunc("/menuitem/{id}", menuItemHandler.DeleteMenuItem).Methods("DELETE")
	api.HandleFunc("/order/{id}", orderHandler.DeleteOrder).Methods("DELETE")
	api.HandleFunc("/itemorder/{id}", itemOrderHandler.DeleteItemOrder).Methods("DELETE")

	// Starting Server
	const serverDoor int = 8080
	ip, err := utils.IPCheck()
	if err != nil {
		log.Println("Ip detection error!")
	}
	apiDoor := fmt.Sprintf(":%d", serverDoor)

	fmt.Println()
	// log.SetOutput(os.Stdout)
	log.Printf("Server is running on port http://%s%s:\n", ip, apiDoor)
	if err := http.ListenAndServe(apiDoor, route); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	// The line below will only be reached if `http.ListenAndServe` returns an error.
	log.Println("Server stopped")
}
