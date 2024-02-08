package main

import (
	"bares_api/bootstrap"
	"bares_api/handlers"
	"bares_api/services"
	"bares_api/store"
	"bares_api/utils"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	const dbName = "BarDB" // Nome do Banco de dados

	// Pega usuário e senha do banco de dados
	connString := utils.CreateConnString(dbName)

	// Start database
	dbStore, err := store.DatabaseOpen(dbName, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer dbStore.DatabaseClose()

	err = dbStore.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Inicializar stores
	userStore := store.NewUser(dbStore.DB)
	menuItemStore := store.NewMenuItem(dbStore.DB)
	orderStore := store.NewOrder(dbStore.DB)
	itemOrderStore := store.NewItemOrder(dbStore.DB)

	// Inicializar services
	userService := services.NewUsuarioService(userStore)
	authService := services.NewAuthservice(userStore)
	menuItemService := services.NewItemMenuService(menuItemStore)
	orderService := services.NewPedidoService(orderStore)
	itemOrderService := services.NewItemPedidoService(itemOrderStore)

	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	menuItemHandler := handlers.NewMenuItemHandler(menuItemService)
	orderHandler := handlers.NewOrderHandler(orderService)
	itemOrderHandler := handlers.NewItemOrderHandler(itemOrderService)

	// Verifica se existe um usuário 'admin' e um 'gerente'.
	if err := bootstrap.CheckAndCreateAdminUser(userService); err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	// Configurar rotas
	route := mux.NewRouter()

	// Rotas públicas
	route.HandleFunc("/login", authHandler.LoginHandlers).Methods("POST")
	route.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// Rotas privadas
	api := route.PathPrefix("/api").Subrouter()
	api.Use(handlers.AuthMiddleware) // Aplica o middleware de autenticação

	// As rotas abaixo requerem autenticação
	api.HandleFunc("/menuitem", menuItemHandler.CreateMenuItem).Methods("POST")
	api.HandleFunc("/order", orderHandler.CreateOrder).Methods("POST")
	api.HandleFunc("/itemorder", menuItemHandler.CreateMenuItem).Methods("POST")

	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/itemmenu/{id}", menuItemHandler.GetMenuItem).Methods("GET")
	api.HandleFunc("/itemmenu", menuItemHandler.GetAllMenuItem).Methods("GET")
	api.HandleFunc("/itemmenu/name/{name}", menuItemHandler.GetMenuItemByNome).Methods("GET")
	api.HandleFunc("/order/{id}", orderHandler.GetOrder).Methods("GET")
	api.HandleFunc("/order/user/{id}", orderHandler.GetOrderByUser).Methods("GET")
	api.HandleFunc("/order", orderHandler.GetPendingOrder).Methods("GET")
	api.HandleFunc("/itemorder/{id}", itemOrderHandler.GetIItemOrder).Methods("GET")

	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/itemmenu/{id}", menuItemHandler.UpdateMenuItem).Methods("PUT")
	api.HandleFunc("/order/{id}", orderHandler.UpdateOrder).Methods("PUT")
	api.HandleFunc("/itemorder/{id}", itemOrderHandler.UpdateItemOrder).Methods("PUT")

	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	api.HandleFunc("/itemmenu/{id}", menuItemHandler.DeleteMenuItem).Methods("DELETE")
	api.HandleFunc("/order/{id}", orderHandler.DeleteOrder).Methods("DELETE")
	api.HandleFunc("/itemorder/{id}", itemOrderHandler.DeleteItemOrder).Methods("DELETE")

	// Iniciar servidor
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	// A linha abaixo só será alcançada se `http.ListenAndServe` retornar um erro.
	log.Println("Server stopped")
}
