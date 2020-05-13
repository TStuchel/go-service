package http

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestHttpFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Http Filters Suite")
}

// Ginkgo BDD tests
var _ = Describe("Http Filters", func() {

	// --
	Describe("No filters and no handler function should return no function in return", func() {

		// GIVEN no filters AND no handler function
		// THEN no function should be returned
		It("should return no function", func() {
			chain := BuildFilterChain(nil, nil)
			Expect(chain).To(BeNil())
		})
	})

	// --
	Describe("A single handler should be the only handler in the chain", func() {
		var (
			handler   http.HandlerFunc
			chain     http.HandlerFunc
			wasCalled bool
		)

		// GIVEN no filters AND a handler function
		BeforeEach(func() {
			handler = func(writer http.ResponseWriter, request *http.Request) {
				wasCalled = true
				return
			}
			chain = BuildFilterChain(nil, handler)
		})

		// WHEN the chain is called
		JustBeforeEach(func() {
			chain(nil, nil)
		})

		It("should call the handler function", func() {
			Expect(wasCalled).To(BeTrue())
		})

	})

	// --
	Describe("Multiple filters should be called in order with the handler being the last function called", func() {
		const filterCount = 5
		var (
			filters   []Filter
			handler   http.HandlerFunc
			chain     http.HandlerFunc
			callIndex []int
		)

		// GIVEN a list of filters AND a handler function
		BeforeEach(func() {

			// Filters
			filters = make([]Filter, filterCount)
			for i := 0; i < len(filters); i++ {
				index := i // required for closure
				filters[index] = func(handlerFunc http.HandlerFunc) http.HandlerFunc {
					return func(writer http.ResponseWriter, request *http.Request) {
						callIndex[index] = index
						handlerFunc(writer, request)
					}
				}
			}

			// Handler
			handler = func(writer http.ResponseWriter, request *http.Request) {
				callIndex[5] = 99
				return
			}

			callIndex = make([]int, filterCount+1)
			chain = BuildFilterChain(filters, handler)
		})

		// WHEN the chain is called
		JustBeforeEach(func() {
			chain(nil, nil)
		})

		It("should call the chain of functions in the correct order", func() {
			for i := 0; i < filterCount; i++ {
				Expect(callIndex[i]).To(Equal(i))
			}
		})

	})

})
