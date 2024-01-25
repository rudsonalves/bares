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
	// Start database
	dbStore, err := store.DatabaseOpen()
	if err != nil {
		log.Fatal(err)
	}
	defer dbStore.DatabaseClose()

	err = dbStore.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Inicializar stores
	usuarioStore := store.NewUsuarioStore(dbStore.DB)
	itemMenuStore := store.NewItensMenuStore(dbStore.DB)
	pedidoStore := store.NewPedidoStore(dbStore.DB)
	itemPedidoStore := store.NewItemPedidoStore(dbStore.DB)

	// Inicializar services
	usuarioService := services.NewUsuarioService(usuarioStore)
	itemMenuService := services.NewItemMenuService(itemMenuStore)
	pedidoService := services.NewPedidoService(pedidoStore)
	itemPedidoService := services.NewItemPedidoService(itemPedidoStore)

	// Inicializar handlers
	usuarioHandle := handlers.NewUsuarioHandler(usuarioService)
	itemMenuHandler := handlers.NewItemMenuHandler(itemMenuService)
	pedidoHandler := handlers.NewPedidoHandler(pedidoService)
	itemPedidoHandler := handlers.NewItemPedidoHandler(itemPedidoService)

	// Configurar rotas
	route := mux.NewRouter()
	route.HandleFunc("/usuarios", usuarioHandle.CreateUsuario).Methods("POST")
	route.HandleFunc("/itemmenu", itemMenuHandler.CreateItemMenu).Methods("POST")
	route.HandleFunc("/pedido", pedidoHandler.CreatePedido).Methods("POST")
	route.HandleFunc("/itempedido", itemMenuHandler.CreateItemMenu).Methods("POST")

	route.HandleFunc("/usuarios/{id}", usuarioHandle.GetUsuario).Methods("GET")
	route.HandleFunc("/itemmenu/{id}", itemMenuHandler.GetItemMenu).Methods("GET")
	route.HandleFunc("/itemmenu", itemMenuHandler.GetAllItemMenu).Methods("GET")
	route.HandleFunc("/itemmenu/name/{name}", itemMenuHandler.GetItemMenuByNome).Methods("GET")
	route.HandleFunc("/pedido/{id}", pedidoHandler.GetPedido).Methods("GET")
	route.HandleFunc("/pedido/usuario/{id}", pedidoHandler.GetPedidosByUsuario).Methods("GET")
	route.HandleFunc("/itempedido/{id}", itemPedidoHandler.GetItemPedido).Methods("GET")

	route.HandleFunc("/usuarios/{id}", usuarioHandle.UpdateUsuario).Methods("PUT")
	route.HandleFunc("/itemmenu/{id}", itemMenuHandler.UpdateItemMenu).Methods("PUT")
	route.HandleFunc("/pedido/{id}", pedidoHandler.UpdatePedido).Methods("PUT")
	route.HandleFunc("/itempedido/{id}", itemPedidoHandler.UpdateItemPedido).Methods("PUT")

	route.HandleFunc("/usuarios/{id}", usuarioHandle.DeleteUsuario).Methods("DELETE")
	route.HandleFunc("/itemmenu/{id}", itemMenuHandler.DeleteItemMenu).Methods("DELETE")
	route.HandleFunc("/pedido/{id}", pedidoHandler.DeletePedido).Methods("DELETE")
	route.HandleFunc("/itempedido/{id}", itemPedidoHandler.DeleteItemPedido).Methods("DELETE")

	// Iniciar servidor
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	// A linha abaixo só será alcançada se `http.ListenAndServe` retornar um erro.
	log.Println("Server stopped")
}
