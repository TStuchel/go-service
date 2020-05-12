package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/gomega"
	"net/http"
	"os"
	"testing"
	"time"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

// GIVEN an HTTP request containing (or not containing) a valid JWT token
// WHEN an attempt is made to extract the JWT token
// THEN the JWT token should be extracted, or an error returned
func TestExtractJwtToken(t *testing.T) {

	// Initialize Gomega assertions
	g := NewGomegaWithT(t)

	// Test testValues
	_ = os.Setenv("JWT_SECRET", "TEST_SECRET")
	goodToken, _ := generateTestToken(time.Now().Unix() + 3600) // 1 hour
	testValues := []jwtHelper{
		{Header: "", Prefix: "", Token: "", Error: errors.New("missing token")},
		{Header: "Authorization", Prefix: "", Token: "", Error: errors.New("missing token")},
		{Header: "Authorization", Prefix: "Basic ", Token: "", Error: errors.New("missing token")},
		{Header: "Authorization", Prefix: "Bearer", Token: "", Error: errors.New("missing token")},
		{Header: "Authorization", Prefix: "Bearer ", Token: "", Error: errors.New("missing token")},
		{Header: "Authorization", Prefix: "Bearer ", Token: "INVALID", Error: errors.New("token contains an invalid number of segments")},
		{Header: "Authorization", Prefix: "Bearer ", Token: goodToken, Error: nil},
	}

	for _, test := range testValues {
		t.Run(fmt.Sprintf("%s %s:%s", test.Header, test.Prefix, test.Token), func(t *testing.T) {
			// GIVEN an HTTP request
			request := &http.Request{Header: http.Header{}}
			if test.Header != "" {
				request.Header.Add(test.Header, fmt.Sprintf("%s%s", test.Prefix, test.Token))
			}

			// WHEN an attempt is made to extract the JWT token
			token, err := ExtractToken(request)

			// THEN the JWT token should be extracted, or an error returned
			if test.Error != nil {
				g.Expect(token).To(BeNil())
				g.Expect(err).NotTo(BeNil())
				g.Expect(err.Error()).To(Equal(test.Error.Error()))
			} else {
				g.Expect(err).To(BeNil())
				g.Expect(token).ToNot(BeNil())
				g.Expect(token.Valid).To(BeTrue())
			}
		})
	}

}

// ------------------------------------------------------ HELPERS ------------------------------------------------------

type jwtHelper struct {
	Header string
	Prefix string
	Token  string
	Error  error
}

func generateTestToken(expiresAt int64) (token string, err error) {

	// Create the claims
	claims := Claims{
		Username: "TEST_USER",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "www.daugherty.com",
		},
	}

	// Create/Sign the token and return
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
