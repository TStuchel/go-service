package auth_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TStuchel/go-service/auth"
	"github.com/TStuchel/go-service/auth/authfakes"
	"github.com/TStuchel/go-service/common"
	. "github.com/TStuchel/go-service/testutils"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestAuthController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

// Ginkgo BDD tests
var _ = Describe("Controller", func() {

	// Per-suite variables
	var (
		router     *mux.Router
		service    *authfakes.FakeService
		controller auth.Controller
	)

	// Per-suite setup
	BeforeEach(func() {
		router = mux.NewRouter()
		service = new(authfakes.FakeService)
		controller = auth.NewAuthController(router, service)
	})

	// --
	Describe("Creating a new auth controller", func() {
		It("should result in a Controller that references the given dependencies", func() {
			Expect(controller).ToNot(BeNil())
		})
		It("should map its routes", func() {
			Expect(router.GetRoute("GetToken")).ToNot(BeNil())
		})
	})

	// --
	Describe("Get a token given valid credentials", func() {
		var (
			req *http.Request
			w   *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			// GIVEN valid credentials
			service.LoginReturns(RandomString(12), nil)

			// WHEN the customer API endpoint is called
			req, _ = http.NewRequest("GET", "/v1/token", nil)
			req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("admin:admin"))))
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
		})
		It("should return an HTTP status code of 200-OK", func() {
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("should contain a newly created token", func() {
			tokenResponse := auth.TokenResponse{}
			_ = json.Unmarshal(w.Body.Bytes(), &tokenResponse)
			Expect(tokenResponse.Token).ToNot(BeEmpty())
		})
	})

	// --
	Describe("Return an error given invalid credentials", func() {
		var (
			req *http.Request
			w   *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			// GIVEN invalid credentials
			service.LoginReturns("", errors.New("invalid credentials"))

			// WHEN the customer API endpoint is called
			req, _ = http.NewRequest("GET", "/v1/token", nil)
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
		})
		It("should return an HTTP status code of 401-Unauthorized", func() {
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
		It("should should contain the expected error", func() {
			response := common.ErrorDTO{}
			_ = json.Unmarshal(w.Body.Bytes(), &response)
			Expect(response.Url).To(Equal("/v1/token"))
			Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
			Expect(response.Message).To(Equal("invalid credentials"))
			Expect(response.Type).To(Equal("UnauthorizedException"))
		})
	})
})
