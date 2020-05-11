package customer

import (
	. "github.com/TStuchel/go-service/common"
	. "github.com/TStuchel/go-service/logging"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

type Controller interface {
	GetCustomer(http.ResponseWriter, *http.Request)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type controllerImpl struct {
	router  *mux.Router
	service Service
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewCustomerController creates and returns a Controller with its handlers registered with the given router.
func NewCustomerController(router *mux.Router, service Service) Controller {

	// Create controller
	controller := controllerImpl{
		router:  router,
		service: service,
	}

	// Register handlers
	router.HandleFunc("/v1/customers/{id}", Logger(controller.GetCustomer)).Methods("GET").Name("GetCustomer")
	router.HandleFunc("/v1/customers", Logger(controller.CreateCustomer)).Methods("POST").Name("CreateCustomer")

	return controller
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetCustomer returns a Customer JSON contract given the customer ID
func (impl controllerImpl) GetCustomer(w http.ResponseWriter, r *http.Request) {

	// Read the URI path variables
	vars := mux.Vars(r)
	customerId := vars["id"]

	// Get the customer
	startTime := time.Now()
	customer, err := impl.service.GetCustomer(customerId) // TODO: Handle Error
	elapsedTime := time.Since(startTime)

	// Build the HTTP response
	w.Header().Set("x-elapsed", elapsedTime.String())

	// Error
	if err != nil {
		HandleBadRequest(w, err)
		return
	}

	// Missing data
	if customer == nil {
		HandleNotFound(w)
		return
	}

	// Translate
	customerDTO := ToContract(customer)

	// Good data, return JSON
	HandleSuccess(w, http.StatusOK, customerDTO)
}

// CreateCustomer creates a new customer with the given data
func (impl controllerImpl) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Boom!")
	}

	log.Printf("Got Request : %s",string(body))
}

