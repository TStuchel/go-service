package customer_test

import (
	"encoding/json"
	"fmt"
	"github.com/TStuchel/go-service/common"
	"github.com/TStuchel/go-service/customer"
	"github.com/TStuchel/go-service/customer/customerfakes"
	"github.com/TStuchel/go-service/testutils"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestCustomerController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Customer Controller Suite")
}

// Ginkgo BDD tests
var _ = Describe("Customer Controller", func() {

	// Per-suite variables
	var (
		router     *mux.Router
		service    *customerfakes.FakeService
		controller customer.Controller
	)

	// Per-suite setup
	BeforeEach(func() {
		router = mux.NewRouter()
		service = new(customerfakes.FakeService)
		controller = customer.NewCustomerController(router, nil, service)
	})

	// --
	Describe("Creating a new customer controller", func() {
		It("should result in a Controller that references the given dependencies", func() {
			Expect(controller).ToNot(BeNil())
		})
		It("should map its routes", func() {
			Expect(router.GetRoute("GetCustomer")).ToNot(BeNil())
		})
	})

	// --
	Describe("Find a customer by its ID", func() {
		customerId := testutils.RandomInt(0, 99999)
		var (
			req *http.Request
			w   *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			// GIVEN a customer with a particular ID is in the system
			service.GetCustomerReturns(&customer.Customer{ID: strconv.Itoa(customerId)}, nil)

			// WHEN the customer API endpoint is called
			req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/customers/%d", customerId), nil)
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
		})
		It("should return an HTTP response code of 200", func() {
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("should contain the Customer in the response", func() {
			customerDTO := new(customer.CustomerDTO)
			_ = json.Unmarshal(w.Body.Bytes(), customerDTO)
			Expect(customerDTO.Id).To(Equal(strconv.Itoa(customerId)))
		})
	})

	// --
	Describe("Respond with NOT_FOUND given an unknown customer ID", func() {
		var (
			req *http.Request
			w   *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			// GIVEN an ID of a customer that is not in the system
			service.GetCustomerReturns(nil, nil)
			customerId := testutils.RandomInt(0, 99999)

			// WHEN the customer API endpoint is called
			req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/customers/%d", customerId), nil)
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
		})
		It("the response HTTP status should be NOT_FOUND", func() {
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
		It("an error body should be empty", func() {
			Expect(w.Body.String()).ToNot(BeNil())
		})
	})

	// --
	Describe("Respond with BAD_REQUEST if given an invalid customer ID", func() {
		var (
			req        *http.Request
			w          *httptest.ResponseRecorder
			customerId int
		)
		BeforeEach(func() {
			// GIVEN an ID of a customer that is not in the system
			customerId = rand.Intn(101)
			busError := common.NewBusinessError(fmt.Sprintf("Invalid customer ID [%d]", customerId))
			service.GetCustomerReturns(nil, busError)

			// WHEN the customer API endpoint is called
			req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/customers/%d", customerId), nil)
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
		})
		It("the response HTTP status should be BAD_REQUEST", func() {
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
		It("an error body should be returned", func() {
			Expect(w.Body.String()).ToNot(BeNil())
		})
		It("should contain the error message", func() {
			busError := new(common.BusinessError)
			_ = json.Unmarshal(w.Body.Bytes(), busError)
			Expect(busError.Data).To(Equal(fmt.Sprintf("Invalid customer ID [%d]", customerId)))
		})
	})
})
