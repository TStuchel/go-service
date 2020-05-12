package customer

import (
	http2 "github.com/TStuchel/go-service/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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
func NewCustomerController(router *mux.Router, filters []http2.Filter, service Service) Controller {

	// Create controller
	controller := controllerImpl{
		router:  router,
		service: service,
	}

	// Register handlers
	router.HandleFunc("/v1/customers/{id}", http2.BuildFilterChain(filters, controller.GetCustomer)).Methods("GET").Name("GetCustomer")
	router.HandleFunc("/v1/customers", http2.BuildFilterChain(filters, controller.CreateCustomer)).Methods("POST").Name("CreateCustomer")

	return controller
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetCustomer returns a Customer JSON contract given the customer ID
func (impl controllerImpl) GetCustomer(w http.ResponseWriter, r *http.Request) {

	// Read the URI path variables
	vars := mux.Vars(r)
	customerId := vars["id"]

	// Get the customer
	customer, err := impl.service.GetCustomer(customerId) // TODO: Handle Error

	// Error
	if err != nil {
		http2.HandleBadRequest(w, err)
		return
	}

	// Missing data
	if customer == nil {
		http2.HandleNotFound(w)
		return
	}

	// Translate
	customerDTO := ToContract(customer)

	// Good data, return JSON
	http2.HandleSuccess(w, http.StatusOK, customerDTO)
}

// CreateCustomer creates a new customer with the given data
func (impl controllerImpl) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Boom!")
	}

	log.Printf("Got Request : %s", string(body))
}
