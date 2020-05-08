package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

// Controller : Root controller interface
type Controller interface {
	GetHealth(http.ResponseWriter, *http.Request)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

// controllerImpl : Root application controller
type controllerImpl struct {
	router *mux.Router
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewAppController : Create and return a new reservation controller with the given dependencies
func NewAppController(router *mux.Router) Controller {

	// Create controller
	controller := controllerImpl{
		router: router,
	}

	// Register handlers
	router.HandleFunc("/health", controller.GetHealth).Methods("GET").Name("GetHealth")

	return controller
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetHealth : Returns the health of this service
func (controllerImpl) GetHealth(w http.ResponseWriter, _ *http.Request) {

	// Build the HTTP response
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	var _, err = w.Write([]byte("Service available"))
	if err != nil {
		log.Print(err)
	}
}
