package main

import (
	"bares_api/bootstrap"
	"bares_api/utils"
	"fmt"
	"log"
	"net/http"
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

	router := bootstrap.SetupRouter(dbStore)

	// Starting Server
	const serverDoor int = 8080
	ip, err := utils.IPCheck()
	if err != nil {
		log.Println("Ip detection error!")
	}
	apiDoor := fmt.Sprintf(":%d", serverDoor)

	fmt.Println()
	// log.SetOutput(os.Stdout)
	log.Printf("Server is running on port http://%s%s\n", ip, apiDoor)
	if err := http.ListenAndServe(apiDoor, router); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	// The line below will only be reached if `http.ListenAndServe` returns an error.
	log.Println("Server stopped")
}
