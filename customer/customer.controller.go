package customer

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/v1/customers/{id}", controller.GetCustomer).Methods("GET").Name("GetCustomer")

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("x-elapsed", elapsedTime.String())

	// Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		enc := json.NewEncoder(w)
		err = enc.Encode(err)
		return
	}

	// Missing data
	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Translate
	customerDTO := ToContract(customer)

	// Good data, return JSON
	enc := json.NewEncoder(w)
	err = enc.Encode(customerDTO)
	if err != nil {
		log.Print(err)
	}
}
