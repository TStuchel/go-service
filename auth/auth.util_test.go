package auth_test

import (
	"encoding/base64"
	"fmt"
	"github.com/TStuchel/go-service/auth"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

// GIVEN an HTTP request containing (or not containing) valid Basic auth credentials
// WHEN the basic authentication credentials are extracted
// THEN the username and password should be extracted, or an error returned
func TestExtractBasicAuthCredentials(t *testing.T) {

	// Initialize Gomega assertions
	g := NewGomegaWithT(t)

	// Possible Authorization credential combinations
	values := [][]string{
		{"", "", "", "", "", "invalid credentials"},
		{"Authorization", "INVALID", "", "", "", "invalid credentials"},
		{"Authorization", "INVALID ", "", "", "", "invalid credentials"},
		{"Authorization", "Basic", "", "", "", "invalid credentials"},
		{"Authorization", "Basic ", "", "", "", "invalid credentials"},
		{"Authorization", "Basic ", "admin:admin", "", "", "invalid credentials"},
		{"Authorization", "Basic ", base64.StdEncoding.EncodeToString([]byte(":admin")), "", "", "invalid credentials"},
		{"Authorization", "Basic ", base64.StdEncoding.EncodeToString([]byte("admin:")), "", "", "invalid credentials"},
		{"Authorization", "Basic ", base64.StdEncoding.EncodeToString([]byte("admin:admin")), "admin", "admin", ""},
	}

	for _, test := range values {
		t.Run(fmt.Sprintf("%s %s:%s", test[0], test[1], test[2]), func(t *testing.T) {

			// Build the request
			request := http.Request{}
			request.Header = http.Header{}
			request.Header.Add(test[0], fmt.Sprintf("%s%s", test[1], test[2]))

			// Extract
			var username, password, err = auth.ExtractBasicAuthCredentials(&request)

			// Verify
			g.Expect(username).To(Equal(test[3]))
			g.Expect(password).To(Equal(test[4]))
			if test[5] != "" {
				g.Expect(err).ToNot(BeNil())
				g.Expect(err.Error()).To(Equal(test[5]))
			} else {
				g.Expect(err).To(BeNil())
			}
		})
	}
}
