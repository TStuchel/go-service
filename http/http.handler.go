package http

import (
	"encoding/json"
	"github.com/TStuchel/go-service/common"
	"log"
	"net/http"
)

// ------------------------------------------------- PUBLIC FUNCTIONS --------------------------------------------------

// HandleSuccess writes to the given ResponseWriter with the given HTTP status code and writes the given structure to it
// as a JSON string.
func HandleSuccess(w http.ResponseWriter, httpStatus int, dto interface{}) {

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)

	// Write the DTO
	if dto != nil {
		enc := json.NewEncoder(w)
		err := enc.Encode(dto)

		// Very bad
		if err != nil {
			log.Print(err)
		}
	}
}

// HandleNotFound writes to the given ResponseWriter with an HTTP status of NOT_FOUND.
func HandleNotFound(w http.ResponseWriter) {

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
}

// HandleBadRequest writes to the given ResponseWriter with an HTTP status of BAD_REQUEST and writes the given error
// as a JSON string.
func HandleBadRequest(w http.ResponseWriter, url string, err error) {

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	// Write the DTO
	if err != nil {
		enc := json.NewEncoder(w)
		err = enc.Encode(common.ErrorDTO{
			Url:        url,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Type:       "BadRequestError",
		})

		// Very bad
		if err != nil {
			log.Print(err)
		}
	}
}

// HandleBadRequest writes to the given ResponseWriter with an HTTP status of UNAUTHORIZED and writes the given error
// as a JSON string.
func HandleUnauthorizedError(w http.ResponseWriter, url string, err error) {

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	// Write the DTO
	if err != nil {
		enc := json.NewEncoder(w)
		err = enc.Encode(common.ErrorDTO{
			Url:        url,
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Type:       "UnauthorizedError",
		})

		// Very bad
		if err != nil {
			log.Print(err)
		}
	}
}
