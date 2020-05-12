package jwt

import (
	"fmt"
	"github.com/TStuchel/go-service/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"os"
	"testing"
	"time"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestJwtFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT Filter Suite")
}

// Ginkgo BDD tests
var _ = Describe("JWT Filter", func() {

	// --
	Describe("A missing or invalid JWT token should return 401-Unauthorized", func() {
		var (
			writer           *testutils.MockHttpResponseWriter
			request          *http.Request
			handler          http.HandlerFunc
			wasHandlerCalled bool
		)

		// GIVEN an HTTP request that does not contain a JWT token
		BeforeEach(func() {
			writer = &testutils.MockHttpResponseWriter{}
			request = &http.Request{}

			// Mock
			handler = func(writer http.ResponseWriter, request *http.Request) {
				wasHandlerCalled = true
			}
		})

		// WHEN the HTTP request is filtered
		JustBeforeEach(func() {
			filter := Filter{}
			filteredHandler := filter.Handle(handler)
			filteredHandler(writer, request)
		})

		It("should not call the handler", func() {
			Expect(wasHandlerCalled).To(BeFalse())
		})
		It("should return a 401-Unauthorized response", func() {
			Expect(writer.StatusCode).To(Equal(http.StatusUnauthorized))
		})

	})

	// --
	Describe("A valid JWT token should call the handler method", func() {
		var (
			writer           *testutils.MockHttpResponseWriter
			request          *http.Request
			handler          http.HandlerFunc
			wasHandlerCalled bool
		)

		// GIVEN an HTTP request that contains a valid JWT token
		BeforeEach(func() {
			writer = &testutils.MockHttpResponseWriter{}
			request = &http.Request{Header: http.Header{}}

			_ = os.Setenv("JWT_SECRET", "TEST_SECRET")
			invalidToken, _ := generateTestToken(time.Now().Unix() + 3600) // 1 hour
			request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", invalidToken))

			// Mock
			handler = func(writer http.ResponseWriter, request *http.Request) {
				wasHandlerCalled = true
			}
		})

		// WHEN the HTTP request is filtered
		JustBeforeEach(func() {
			filter := Filter{}
			filteredHandler := filter.Handle(handler)
			filteredHandler(writer, request)
		})

		It("should call the handler method", func() {
			Expect(wasHandlerCalled).To(BeTrue())
		})

	})
})
