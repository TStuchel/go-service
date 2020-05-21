package http

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestPerformanceFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Performance Filter Suite")
}

// Ginkgo BDD tests
var _ = Describe("Performance Filter", func() {

	// --
	Describe("The performance filter should add a elapsed time header to the response", func() {
		var (
			writer  *MockHttpResponseWriter
			request *http.Request
		)

		// GIVEN an HTTP request
		BeforeEach(func() {
			writer = &MockHttpResponseWriter{}
			request = &http.Request{Header: http.Header{}}
		})

		// WHEN the HTTP request is filtered
		JustBeforeEach(func() {
			filteredHandler := PerformanceFilter(func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Add("SomeOtherHeader", "value")
				return
			})
			filteredHandler(writer, request)
		})

		// THEN
		It("should keep any added headers", func() {
			Expect(writer.Header().Get("SomeOtherHeader")).To(Equal("value"))
		})
		It("should add a elapsed time header to the response", func() {
			Expect(writer.Header().Get("X-Elapsed")).ToNot(BeNil())
		})

	})
})
