package main

import (
	"bares_api/handlers"
	"bares_api/services"
	"bares_api/store"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	const dbName = "BarDB" // Nome do Banco de dados

	// Start database
	dbStore, err := store.DatabaseOpen(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer dbStore.DatabaseClose()

	err = dbStore.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Inicializar stores
	usuarioStore := store.NewUser(dbStore.DB)
	itemMenuStore := store.NewMenuItem(dbStore.DB)
	pedidoStore := store.NewOrder(dbStore.DB)
	itemPedidoStore := store.NewItemOrder(dbStore.DB)

	// Inicializar services
	usuarioService := services.NewUsuarioService(usuarioStore)
	authService := services.NewAuthservice(usuarioStore)
	itemMenuService := services.NewItemMenuService(itemMenuStore)
	pedidoService := services.NewPedidoService(pedidoStore)
	itemPedidoService := services.NewItemPedidoService(itemPedidoStore)

	// Inicializar handlers
	usuarioHandler := handlers.NewUserHandler(usuarioService)
	authHandler := handlers.NewAuthHandler(authService)
	itemMenuHandler := handlers.NewMenuItemHandler(itemMenuService)
	pedidoHandler := handlers.NewOrderHandler(pedidoService)
	itemPedidoHandler := handlers.NewItemOrderHandler(itemPedidoService)

	// Configurar rotas
	route := mux.NewRouter()

	// Rotas públicas
	route.HandleFunc("/login", authHandler.LoginHandlers).Methods("POST")
	route.HandleFunc("/usuarios", usuarioHandler.CreateUser).Methods("POST")

	// Rotas privadas
	api := route.PathPrefix("/api").Subrouter()
	api.Use(handlers.AuthMiddleware) // Aplica o middleware de autenticação

	// As rotas abaixo requerem autenticação
	api.HandleFunc("/itemmenu", itemMenuHandler.CreateMenuItem).Methods("POST")
	api.HandleFunc("/pedido", pedidoHandler.CreateOrder).Methods("POST")
	api.HandleFunc("/itempedido", itemMenuHandler.CreateMenuItem).Methods("POST")

	api.HandleFunc("/usuarios/{id}", usuarioHandler.GetUser).Methods("GET")
	api.HandleFunc("/itemmenu/{id}", itemMenuHandler.GetMenuItem).Methods("GET")
	api.HandleFunc("/itemmenu", itemMenuHandler.GetAllMenuItem).Methods("GET")
	api.HandleFunc("/itemmenu/name/{name}", itemMenuHandler.GetMenuItemByNome).Methods("GET")
	api.HandleFunc("/pedido/{id}", pedidoHandler.GetOrder).Methods("GET")
	api.HandleFunc("/pedido/usuario/{id}", pedidoHandler.GetOrderByUser).Methods("GET")
	api.HandleFunc("/pedido", pedidoHandler.GetPendingOrder).Methods("GET")
	api.HandleFunc("/itempedido/{id}", itemPedidoHandler.GetIItemOrder).Methods("GET")

	api.HandleFunc("/usuarios/{id}", usuarioHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/itemmenu/{id}", itemMenuHandler.UpdateMenuItem).Methods("PUT")
	api.HandleFunc("/pedido/{id}", pedidoHandler.UpdateOrder).Methods("PUT")
	api.HandleFunc("/itempedido/{id}", itemPedidoHandler.UpdateItemOrder).Methods("PUT")

	api.HandleFunc("/usuarios/{id}", usuarioHandler.DeleteUser).Methods("DELETE")
	api.HandleFunc("/itemmenu/{id}", itemMenuHandler.DeleteMenuItem).Methods("DELETE")
	api.HandleFunc("/pedido/{id}", pedidoHandler.DeleteOrder).Methods("DELETE")
	api.HandleFunc("/itempedido/{id}", itemPedidoHandler.DeleteItemOrder).Methods("DELETE")

	// Iniciar servidor
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	// A linha abaixo só será alcançada se `http.ListenAndServe` retornar um erro.
	log.Println("Server stopped")
}
