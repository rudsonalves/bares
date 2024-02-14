package utils

import (
	"bares_api/store"
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

// Create DataBase Conn creates the connection to the database.
func CreateDataBaseConn(dbName string) *store.DatabaseStore {
	reader := bufio.NewReader(os.Stdin)
	var connectionString string
	var dbStore *store.DatabaseStore

	for {
		fmt.Println("===================================")
		fmt.Println("=            Bares API            =")
		fmt.Println("=          Version 1.0.0          =")
		fmt.Println("===================================")
		fmt.Printf("\nUsername and password to access the database %s.\n", dbName)

		// Database user name
		fmt.Printf("Enter the name: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the name, try again.")
			continue
		}
		name = TrimSpaceLB(name)

		// Database user password
		fmt.Printf("Enter password: ")
		passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password, try again.")
			continue
		}
		password := string(passwordBytes)
		fmt.Printf("\n\n")

		// create connection string
		connectionString = fmt.Sprintf("%s:%s@tcp(localhost:3306)/", name, password)

		dbStore, err = store.DatabaseOpen(dbName, connectionString)
		if err != nil {
			log.Println("Failed to connect to the database:", err)
			fmt.Printf("Trying again...\n\n")
			continue
		}
		break
	}
	println("Database Connected!")

	return dbStore
}
