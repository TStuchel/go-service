package jwt

import (
	http2 "github.com/TStuchel/go-service/http"
	"net/http"
)

// ------------------------------------------------- PUBLIC FUNCTIONS --------------------------------------------------

// This filter verifies the presence of a valid Authorization: Bearer JWT token. If a valid JWT token does not exist
// then this filter shortcuts further processing and returns 401-Unauthorized.
func Filter(handlerFunc http.HandlerFunc) http.HandlerFunc {
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
