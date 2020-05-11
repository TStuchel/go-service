package customer

import (
	. "github.com/TStuchel/go-service/common"
	"github.com/gorilla/mux"
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
