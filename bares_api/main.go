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
  usuarioStore := store.NewUsuario(dbStore.DB)
  itemMenuStore := store.NewItensMenu(dbStore.DB)
  pedidoStore := store.NewPedido(dbStore.DB)
  itemPedidoStore := store.NewItemPedido(dbStore.DB)

  // Inicializar services
  usuarioService := services.NewUsuarioService(usuarioStore)
  authService := services.NewAuthservice(usuarioStore)
  itemMenuService := services.NewItemMenuService(itemMenuStore)
  pedidoService := services.NewPedidoService(pedidoStore)
  itemPedidoService := services.NewItemPedidoService(itemPedidoStore)

  // Inicializar handlers
  usuarioHandler := handlers.NewUsuarioHandler(usuarioService)
  authHandler := handlers.NewAuthHandler(authService)
  itemMenuHandler := handlers.NewItemMenuHandler(itemMenuService)
  pedidoHandler := handlers.NewPedidoHandler(pedidoService)
  itemPedidoHandler := handlers.NewItemPedidoHandler(itemPedidoService)

  // Configurar rotas
  route := mux.NewRouter()

  // Rotas públicas
  route.HandleFunc("/login", authHandler.LoginHandlers).Methods("POST")
  route.HandleFunc("/usuarios", usuarioHandler.CreateUsuario).Methods("POST")

  // Rotas privadas
  api := route.PathPrefix("/api").Subrouter()
  api.Use(handlers.AuthMiddleware) // Aplica o middleware de autenticação

  // As rotas abaixo requerem autenticação
  api.HandleFunc("/itemmenu", itemMenuHandler.CreateItemMenu).Methods("POST")
  api.HandleFunc("/pedido", pedidoHandler.CreatePedido).Methods("POST")
  api.HandleFunc("/itempedido", itemMenuHandler.CreateItemMenu).Methods("POST")

  api.HandleFunc("/usuarios/{id}", usuarioHandler.GetUsuario).Methods("GET")
  api.HandleFunc("/itemmenu/{id}", itemMenuHandler.GetItemMenu).Methods("GET")
  api.HandleFunc("/itemmenu", itemMenuHandler.GetAllItemMenu).Methods("GET")
  api.HandleFunc("/itemmenu/name/{name}", itemMenuHandler.GetItemMenuByNome).Methods("GET")
  api.HandleFunc("/pedido/{id}", pedidoHandler.GetPedido).Methods("GET")
  api.HandleFunc("/pedido/usuario/{id}", pedidoHandler.GetPedidosByUsuario).Methods("GET")
  api.HandleFunc("/pedido", pedidoHandler.GetPedidosPending).Methods("GET")
  api.HandleFunc("/itempedido/{id}", itemPedidoHandler.GetItemPedido).Methods("GET")

  api.HandleFunc("/usuarios/{id}", usuarioHandler.UpdateUsuario).Methods("PUT")
  api.HandleFunc("/itemmenu/{id}", itemMenuHandler.UpdateItemMenu).Methods("PUT")
  api.HandleFunc("/pedido/{id}", pedidoHandler.UpdatePedido).Methods("PUT")
  api.HandleFunc("/itempedido/{id}", itemPedidoHandler.UpdateItemPedido).Methods("PUT")

  api.HandleFunc("/usuarios/{id}", usuarioHandler.DeleteUsuario).Methods("DELETE")
  api.HandleFunc("/itemmenu/{id}", itemMenuHandler.DeleteItemMenu).Methods("DELETE")
  api.HandleFunc("/pedido/{id}", pedidoHandler.DeletePedido).Methods("DELETE")
  api.HandleFunc("/itempedido/{id}", itemPedidoHandler.DeleteItemPedido).Methods("DELETE")

  // Iniciar servidor
  log.Println("Server is running on port 8080")
  if err := http.ListenAndServe(":8080", route); err != nil {
    log.Fatal("Error starting server: ", err)
  }

  // A linha abaixo só será alcançada se `http.ListenAndServe` retornar um erro.
  log.Println("Server stopped")
}
