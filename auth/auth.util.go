package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/TStuchel/go-service/common"
	"log"
	"net/http"
	"strings"
)

// ExtractBasicAuthCredentials extracts the username and password from the Basic auth header.
func ExtractBasicAuthCredentials(r *http.Request) (username string, password string, err error) {

	// Must have the header
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", "", errors.New("invalid credentials")
	}

	// Must have two parts
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", "", errors.New("invalid credentials")
	}

	// Must be Basic Auth
	if parts[0] != "Basic" {
		return "", "", errors.New("invalid credentials")
	}

	// Must be base64
	var bytes, berr = base64.StdEncoding.DecodeString(parts[1])
	if berr != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Must be username:password
	credentials := string(bytes)
	credParts := strings.Split(credentials, ":")
	if len(credParts) != 2 || credParts[0] == "" || credParts[1] == "" {
		return "", "", errors.New("invalid credentials")
	}

	// Username / Password
	return credParts[0], credParts[1], nil
}

// HandleUnauthorizedError writes a 401 Unauthorized error to the given writer.
func HandleUnauthorizedError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	enc := json.NewEncoder(w)
	err = enc.Encode(common.ErrorDTO{
		Url:        "/v1/token",
		StatusCode: http.StatusUnauthorized,
		Message:    err.Error(),
		Type:       "UnauthorizedException",
	})
	if err != nil {
		log.Print(err)
	}
}
