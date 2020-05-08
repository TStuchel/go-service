package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ------------------------------------------ Interfaces ------------------------------------------

// AppController : Root controller interface
type AppController interface {
	GetHealth(http.ResponseWriter, *http.Request)
}

// --------------------------------------------------- Implementation --------------------------------------------------

// AppControllerImpl : Root application controller
type AppControllerImpl struct {
	router *mux.Router
}

// ---------------------------------------------------- Constructor ----------------------------------------------------

// NewAppController : Create and return a new reservation controller with the given dependencies
func NewAppController(router *mux.Router) AppController {

	// Create controller
	controller := AppControllerImpl{
		router: router,
	}

	// Register handlers
	router.HandleFunc("/health", controller.GetHealth).Methods("GET").Name("GetHealth")

	return controller
}

// ------------------------------------------------------ Methods ------------------------------------------------------

// GetHealth : Returns the health of this service
func (impl AppControllerImpl) GetHealth(w http.ResponseWriter, _ *http.Request) {

	// Build the HTTP response
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	_, _ = w.Write([]byte("Service available"))
}
