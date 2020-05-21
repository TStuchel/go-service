package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

// ------------------------------------------------- PUBLIC FUNCTIONS --------------------------------------------------

// ExtractToken extracts the JWT token from the given Request.
func ExtractToken(request *http.Request) (*jwt.Token, error) {

	// Must have a Authorization header
	bearer := request.Header.Get("Authorization")
	if bearer == "" {
		return nil, errors.New("missing token")
	}

	// Must have a 2 part Authorization header
	parts := strings.Split(bearer, " ")
	if len(parts) < 2 {
		return nil, errors.New("missing token")
	}

	// Must be a Bearer token
	if parts[0] != "Bearer" {
		return nil, errors.New("missing token")
	}

	// Must have a token
	if parts[1] == "" {
		return nil, errors.New("missing token")
	}

	// Extract the JWT token
	token, err := jwt.ParseWithClaims(
		parts[1],
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	// Return the token
	return token, err
}
