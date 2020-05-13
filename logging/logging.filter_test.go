package logging

import (
	"bytes"
	"github.com/TStuchel/go-service/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/url"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestLoggingFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logging Filter Suite")
}

// Ginkgo BDD tests
var _ = Describe("Logging Filter", func() {

	// --
	Describe("The logging filter should log before and after calling the handler function", func() {
		var (
			writer           *testutils.MockHttpResponseWriter
			request          *http.Request
			handler          http.HandlerFunc
			wasHandlerCalled bool
		)

		// GIVEN an HTTP request
		BeforeEach(func() {

			writer = &testutils.MockHttpResponseWriter{}
			request = &http.Request{
				URL:  &url.URL{Path: "/test"},
				Body: &testutils.MockReadCloser{Reader: bytes.NewReader([]byte("Some Text"))},
			}

			// Mock
			handler = func(writer http.ResponseWriter, request *http.Request) {
				wasHandlerCalled = true
			}
		})

		// WHEN the HTTP request is filtered
		JustBeforeEach(func() {
			filteredHandler := Filter(handler)
			filteredHandler(writer, request)
		})

		It("should call the handler", func() {
			Expect(wasHandlerCalled).To(BeTrue())
		})

	})

})
