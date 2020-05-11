package common

import (
	"encoding/json"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestHttpHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Http Handlers Suite")
}

// Ginkgo BDD tests
var _ = Describe("Http Handlers", func() {

	Context("HandleUnauthorizedError", func() {
		var (
			err      error
			w        MockHttpResponseWriter
			errorDTO *ErrorDTO
		)

		// GIVEN an (unauthorized) error
		BeforeEach(func() {
			err = errors.New("invalid credentials")
			w = MockHttpResponseWriter{}
		})

		// WHEN the error is handled
		JustBeforeEach(func() {
			HandleUnauthorizedError(&w, "/v1/token", err)
			errorDTO = &ErrorDTO{}
			_ = json.Unmarshal(w.bytes, errorDTO)
		})
		It("should contain the expected error data", func() {
			Expect(errorDTO.StatusCode).To(Equal(http.StatusUnauthorized))
			Expect(errorDTO.Url).To(Equal("/v1/token"))
			Expect(errorDTO.Message).To(Equal("invalid credentials"))
			Expect(errorDTO.Type).To(Equal("UnauthorizedException"))
		})

	})
})

// ------------------------------------------------------ HELPER -------------------------------------------------------

type MockHttpResponseWriter struct {
	statusCode int
	header     http.Header
	bytes      []byte
}

func (impl *MockHttpResponseWriter) Header() http.Header {
	if impl.header == nil {
		impl.header = http.Header{}
	}
	return impl.header
}

func (impl *MockHttpResponseWriter) Write(bytes []byte) (int, error) {
	impl.bytes = bytes
	return 0, nil
}

func (impl *MockHttpResponseWriter) WriteHeader(statusCode int) {
	impl.statusCode = statusCode
}
