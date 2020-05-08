package customer

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// ------------------------------------------ Interfaces ------------------------------------------

//go:generate counterfeiter . CustomerController
type CustomerController interface {
	GetCustomer(http.ResponseWriter, *http.Request)
}

// --------------------------------------------------- Implementation --------------------------------------------------

type CustomerControllerImpl struct {
	router             *mux.Router
	customerService    CustomerService
}

// ---------------------------------------------------- Constructor ----------------------------------------------------

func NewCustomerController(router *mux.Router, customerService CustomerService) CustomerController {

	// Create controller
	controller := CustomerControllerImpl{
		router:             router,
		customerService:    customerService,
	}

	// Register handlers
	router.HandleFunc("/v1/customers/{id}", controller.GetCustomer).Methods("GET").Name("GetCustomer")

	return controller
}

// ------------------------------------------------------ Methods ------------------------------------------------------

// GetCustomer : Return a Customer JSON contract given the customer ID
func (impl CustomerControllerImpl) GetCustomer(w http.ResponseWriter, r *http.Request) {

	// Read the URI path variables
	vars := mux.Vars(r)
	customerId := vars["id"]

	// Get the customer
	startTime := time.Now()
	customer, err := impl.customerService.GetCustomer(customerId) // TODO: Handle Error
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
		panic(err)
	}
}
