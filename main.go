package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/TStuchel/go-service/app"
	"github.com/TStuchel/go-service/auth"
	"github.com/TStuchel/go-service/auth/jwt"
	"github.com/TStuchel/go-service/customer"
	http2 "github.com/TStuchel/go-service/http"
	"github.com/TStuchel/go-service/logging"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Remove this and read from environment variables
func setEnvironment() {
	_ = os.Setenv("APP_PORT", ":8081")
	_ = os.Setenv("JWT_SECRET", "SuperSecretTokenToSignJWT") // TODO: Injected into environment
	_ = os.Setenv("DB_URI", "mongodb://localhost:27017")
	_ = os.Setenv("DB_NAME", "go-service")
}

func main() {
	setEnvironment() // Comment out for real

	// Initialize
	log.Printf("Server starting...")
	router := mux.NewRouter()

	// Connect to the database
	db := NewMongoDB()

	// Basic filters
	baseFilters := []http2.Filter{
		http2.PerformanceFilter,
	}

	// Filters for controllers that are JWT
	jwtFilters := []http2.Filter{
		http2.PerformanceFilter,
		logging.Filter,
		jwt.Filter,
	}

	// Initialize modules
	app.NewAppController(router, baseFilters)

	// Auth
	authService := auth.NewAuthService()
	auth.NewAuthController(router, baseFilters, authService)

	// Customers
	customerRepository := customer.NewCustomerRepository(db)
	customerService := customer.NewCustomerService(customerRepository)
	customer.NewCustomerController(router, jwtFilters, customerService)

	// Create the web server
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0%s", os.Getenv("APP_PORT")),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Run the server in a non-blocking co-routine
	go func() {
		log.Printf("Starting web service on %s...", os.Getenv("APP_PORT"))
		//log.Fatal(srv.ListenAndServeTLS(os.Getenv("APP_PORT"), certFile, keyFile, router))
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for a kill signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // block

	// Shut down the server
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Wait for any connections to finish (no more than timeout) and shut down
	_ = srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}

// NewMongoDB connects to and returns a connection to the database
func NewMongoDB() *mongo.Database {

	// Initialize database
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Unable to open database connection at [%s]. Error: %s", os.Getenv("DB_URI"), err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Unable to ping database connection at [%s]. Error: %s", os.Getenv("DB_URI"), err)
	}

	// Return the database reference
	return client.Database(os.Getenv("DB_NAME"))
}
