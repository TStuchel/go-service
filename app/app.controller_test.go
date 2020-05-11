package app

import (
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestAppController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

// Ginkgo BDD tests
var _ = Describe("Controller", func() {

	// Per-suite variables
	var (
		router *mux.Router
	)

	// Per-suite setup
	BeforeEach(func() {
		router = mux.NewRouter()
	})

	Describe("Creating the application controller", func() {
		Context("with all of its dependencies provided", func() {
			var (
				controller Controller
			)
			BeforeEach(func() {
				controller = NewAppController(router)
			})
			It("should result in a Controller that references the given dependencies", func() {
				Expect(controller).ToNot(BeNil())
			})
			It("should map its routes", func() {
				Expect(router.GetRoute("GetHealth")).ToNot(BeNil())
			})
		})
	})

	Describe("Calling the /health endpoint", func() {
		var (
			req *http.Request
			w   *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			NewAppController(router)
		})
		Context("with a GET request", func() {
			BeforeEach(func() {
				req, _ = http.NewRequest("GET", "/health", nil)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})
			It("should return an HTTP response code of 200", func() {
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})
			It("should contain the response body 'Healthy'", func() {
				Expect(w.Body.String()).To(Equal("Service available"))
			})
		})
	})
})
