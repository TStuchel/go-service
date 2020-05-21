package auth

import (
	http2 "github.com/TStuchel/go-service/pkg/http"
	"github.com/gorilla/mux"
	"net/http"
)

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

type Controller interface {
	GetToken(http.ResponseWriter, *http.Request)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type controllerImpl struct {
	router  *mux.Router
	service Service
}

type TokenResponse struct {
	Token string `json:"token"`
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewAuthController creates and returns an Controller with its handlers registered with the given router.
func NewAuthController(router *mux.Router, filters []http2.Filter, service Service) Controller {

	// Create controller
	controller := controllerImpl{
		router:  router,
		service: service,
	}

	// Register handlers
	router.HandleFunc("/v1/token", http2.BuildFilterChain(filters, controller.GetToken)).Methods("GET").Name("GetToken")

	return controller
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetToken returns a JWT token in the response given valid Basic auth credentials in the request.
func (impl controllerImpl) GetToken(w http.ResponseWriter, r *http.Request) {

	// Get the Basic Auth credentials
	username, password, authErr := ExtractBasicAuthCredentials(r)
	if authErr != nil {
		http2.HandleUnauthorizedError(w, r.URL.Path, authErr)
		return
	}

	// Get the token
	var token, err = impl.service.Login(username, password)
	if err != nil {
		http2.HandleUnauthorizedError(w, r.URL.Path, err)
		return
	}

	// Return token
	http2.HandleSuccess(w, http.StatusOK, TokenResponse{Token: token})
}
