package jwt

import "github.com/dgrijalva/jwt-go"

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

// Claims contains the JWT claims data contained in the JWT token.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
