package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

// ---------------------------------------------------- INTERFACES -----------------------------------------------------

type authClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//go:generate counterfeiter . Service
type Service interface {
	Login(username string, password string) (token string, err error)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type serviceImpl struct {
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

func NewAuthService() Service {
	return &serviceImpl{}
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// Login verifies the given credentials and returns a JWT token if valid.
func (serviceImpl) Login(username string, password string) (token string, err error) {

	// TODO: Authenticate User
	if username != "admin" || password != "admin" {
		return "", errors.New("invalid credentials")
	}

	// Return the token
	return generateToken(username)
}

func generateToken(username string) (token string, err error) {
	// Create the claims
	claims := authClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "www.daugherty.com",
		},
	}

	// Create/Sign the token and return
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
