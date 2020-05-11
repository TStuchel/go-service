package auth_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TStuchel/go-service/auth"
	"github.com/TStuchel/go-service/common"
	"net/http"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

// GIVEN an HTTP request containing (or not containing) valid Basic auth credentials
// WHEN the basic authentication credentials are extracted
// THEN the username and password should be extracted, or an error returned
func TestExtractBasicAuthCredentials(t *testing.T) {

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
			if username != test[3] {
				t.Errorf("Username: Expected [%s], Got [%s]", username, test[3])
			}
			if password != test[4] {
				t.Errorf("Password: Expected [%s], Got [%s]", username, test[4])
			}
			if err != nil && err.Error() != test[5] {
				t.Errorf("Error: Expected [%s], Actual [%s]", username, test[5])
			}
		})
	}
}

// GIVEN an (unauthorized) error
// WHEN the error is handled
// THEN the response should have an HTTP status code of 401-Unauthorized
// AND the response body should contain the expected error
func TestHandleUnauthorizedError(t *testing.T) {

	// GIVEN an (unauthorized) error
	err := errors.New("invalid credentials")
	w := MockHttpResponseWriter{}

	// WHEN the error is handled
	auth.HandleUnauthorizedError(&w, err)
	errorDTO := common.ErrorDTO{}
	_ = json.Unmarshal(w.bytes, &errorDTO)

	// THEN the response should have an HTTP status code of 401-Unauthorized
	if errorDTO.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode : Expected [%d], Actual [%d]", http.StatusUnauthorized, errorDTO.StatusCode)
	}

	// AND the response body should contain the expected error data
	if errorDTO.Url != "/v1/token" {
		t.Errorf("Url : Expected [%s], Actual [%s]", "/v1/token", errorDTO.Url)
	}
	if errorDTO.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode : Expected [%d], Actual [%d]", http.StatusUnauthorized, errorDTO.StatusCode)
	}
	if errorDTO.Message != "invalid credentials" {
		t.Errorf("Message : Expected [%s], Actual [%s]", "invalid credentials", errorDTO.Message)
	}
	if errorDTO.Type != "UnauthorizedException" {
		t.Errorf("Type : Expected [%s], Actual [%s]", "UnauthorizedException", errorDTO.Type)
	}
}

// ------------------------------------------------------ HELPER -------------------------------------------------------

type MockHttpResponseWriter struct {
	statusCode int
	bytes []byte
}

func (impl *MockHttpResponseWriter) Header() http.Header {
	panic("implement me")
}

func (impl *MockHttpResponseWriter) Write(bytes []byte) (int, error) {
	impl.bytes = bytes
	return 0, nil
}

func (impl *MockHttpResponseWriter) WriteHeader(statusCode int) {
	impl.statusCode = statusCode
}

