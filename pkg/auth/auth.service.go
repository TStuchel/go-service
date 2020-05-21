package auth

import (
	"errors"
	jwt2 "github.com/TStuchel/go-service/pkg/auth/jwt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// ---------------------------------------------------- INTERFACES -----------------------------------------------------

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
	claims := jwt2.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600, // 1 hour
			Issuer:    "www.daugherty.com",
		},
	}

	// Create/Sign the token and return
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
