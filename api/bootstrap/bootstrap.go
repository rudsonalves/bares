package bootstrap

import (
	"bares_api/handlers"
	"bares_api/models"
	"bares_api/services"
	"bares_api/store"
	"bares_api/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/term"

	_ "github.com/go-sql-driver/mysql"
)

func SetupRouter(dbStore *store.DatabaseStore) *mux.Router {
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
	if err := CheckAndCreateAdminUser(userService); err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	// Configure routes
	router := mux.NewRouter()

	// Public Routes
	router.HandleFunc("/login", authHandler.LoginHandlers).Methods("POST")

	// Private Routes
	api := router.PathPrefix("/api").Subrouter()
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

	return router
}

// CheckAndCreateAdminUser checks if an Admin user exists. If it does not exist
// an admin user must be created.
func CheckAndCreateAdminUser(userService *services.UserService) error {
	exists, err := userService.CheckIfAdminExists()
	if err != nil {
		log.Printf("Error checking for administrators: %v", err)
		return err
	}

	if exists {
		fmt.Println("Admin user verified!")
		return nil
	}

	// Create admin user
	fmt.Printf("\n\n\n----------------------------------------------------------\n")
	fmt.Printf("\n\n\nNo administrators were found on this system.\n")
	fmt.Println("You must create an Administrator account to continue.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("==================================")
		fmt.Println("=        Enter information       =")
		fmt.Println("=            user Admin          =")
		fmt.Println("==================================")

		var name string
		fmt.Printf("\nAdmin user name: ")
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the name, try again.")
			continue
		}
		name = utils.TrimSpaceLB(name)

		var email string
		fmt.Printf("Admin email adreess: ")
		email, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading email, please try again.")
			continue
		}
		email = utils.TrimSpaceLB(email)

		var password string
		var checkPassword string
		for {
			fmt.Println("Use a password of at least 8 characters, with letters and numbers")
			fmt.Printf("Enter a password: ")
			passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error reading password, try again.")
				continue
			}
			password = string(passwordBytes)
			fmt.Println()

			fmt.Printf("Password check: ")
			passwordBytes, err = term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error reading password, try again.")
				continue
			}
			checkPassword = string(passwordBytes)
			fmt.Println()

			if password != checkPassword {
				fmt.Printf("\nMust-have password! Try again.\n\n")
				continue
			}

			passStrength := utils.EvaluatePasswordStrength(password)
			if passStrength.Score < 6 {
				fmt.Printf("\n\nWeak password! Use a strong password.\n")
				continue
			}

			break
		}

		user := models.User{
			Name:         name,
			Email:        email,
			PasswordHash: password,
			Role:         models.Admin,
		}
		fmt.Printf("\n\nCreate Admin User:\n")
		fmt.Println(user)
		fmt.Println("Confirm the creation of Admin (Y/n): ")

		var ans string
		ans, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation, please try again.")
			continue
		}
		ans = utils.TrimSpaceLB(ans)

		if strings.EqualFold(ans, "y") || len(ans) == 0 {
			err = userService.CreateUser(&user)
			if err != nil {
				log.Printf("Failed to create admin user: %v", err)
				return err
			}
			fmt.Println("Administrator user created successfully!")
			break
		} else if strings.EqualFold(ans, "n") {
			fmt.Println("Administrator user creation cancelled.")
			continue
		} else {
			fmt.Println("Invalid response. Please answer with 'Y' for yes or 'N' for no.")
		}
	}

	return nil
}
