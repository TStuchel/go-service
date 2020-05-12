package jwt

import (
	http2 "github.com/TStuchel/go-service/http"
	"net/http"
)

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

// This filter verifies the presence of a valid Authorization: Bearer JWT token. If a valid JWT token does not exist
// then this filter shortcuts further processing and returns 401-Unauthorized.
type Filter struct {
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

func (Filter) Handle(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// Verify the JWT token
		_, err := ExtractToken(request)
		if err != nil {
			http2.HandleUnauthorizedError(writer, request.RequestURI, err)
			return
		}

		// Call the next filter
		handlerFunc(writer, request)
	}
}
