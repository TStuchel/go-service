package main

import (
	"github.com/TStuchel/go-service/customer"
	"log"
	"net/http"
	"os"

	"github.com/TStuchel/go-service/app"
	"github.com/gorilla/mux"
)

// Remove this and read from environment variables
func setEnvironment() {
	os.Setenv("APP_PORT", ":8081")
}

func main() {
	setEnvironment()

	// Initialize
	log.Printf("Server starting...")

	// Initialize modules
	router := mux.NewRouter()
	app.NewAppController(router)

	// Customers
	customerService := customer.NewCustomerService()
	customer.NewCustomerController(router, customerService)

	// Start web server
	log.Printf("Starting web service on %s...", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), router))
	//log.Fatal(http.ListenAndServeTLS(os.Getenv("APP_PORT"), certFile, keyFile, router))
}
