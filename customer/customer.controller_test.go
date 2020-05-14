package customer_test

import (
	"bytes"
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
	Context("GetCustomer", func() {

		// --
		Describe("Find a customer by its ID", func() {
			var (
				req        *http.Request
				w          *httptest.ResponseRecorder
				customerId int
			)

			// GIVEN a customer with a particular ID is in the system
			BeforeEach(func() {
				customerId = testutils.RandomInt(0, 99999)

				// Mock
				service.GetCustomerReturns(&customer.Customer{ID: strconv.Itoa(customerId)}, nil)
			})

			// WHEN the customer API endpoint is called
			JustBeforeEach(func() {
				req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/customers/%d", customerId), nil)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})

			// THEN
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
				req        *http.Request
				w          *httptest.ResponseRecorder
				customerId int
			)

			// GIVEN an ID of a customer that is not in the system
			BeforeEach(func() {
				customerId = testutils.RandomInt(0, 99999)

				// Mock
				service.GetCustomerReturns(nil, nil)
			})

			// WHEN the customer API endpoint is called
			JustBeforeEach(func() {
				req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/customers/%d", customerId), nil)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})

			// THEN
			It("should return an HTTP response status of NOT_FOUND", func() {
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
			It("should return an empty error body", func() {
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

			// GIVEN an ID of a customer that is not in the system
			BeforeEach(func() {
				customerId = rand.Intn(101)
				busError := common.NewBusinessError(fmt.Sprintf("Invalid customer ID [%d]", customerId))

				// Mock
				service.GetCustomerReturns(nil, busError)
			})

			// WHEN the customer API endpoint is called
			JustBeforeEach(func() {
				req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/customers/%d", customerId), nil)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})

			// THEN
			It("should return a response HTTP status of BAD_REQUEST", func() {
				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
			It("should return an an error body ", func() {
				Expect(w.Body.String()).ToNot(BeNil())
			})
			It("should contain the error message", func() {
				busError := new(common.BusinessError)
				_ = json.Unmarshal(w.Body.Bytes(), busError)
				Expect(busError.Data).To(Equal(fmt.Sprintf("Invalid customer ID [%d]", customerId)))
			})

		})
	})

	// --
	Context("CreateCustomer", func() {

		// --
		Describe("Calling the creation endpoint with a valid DTO should create the customer", func() {
			var (
				req         *http.Request
				w           *httptest.ResponseRecorder
				customerDto customer.CustomerDTO
			)

			// GIVEN the information for a new customer
			BeforeEach(func() {
				customerDto = customer.CustomerDTO{}
				testutils.PopulateTestData(&customerDto)

				// Mock
				customerEntity := customer.ToEntity(customerDto)
				service.CreateCustomerReturns(&customerEntity, nil)
			})

			// WHEN the customer creation endpoint is called
			JustBeforeEach(func() {
				customerDto.Id = ""
				jsonBytes, _ := json.Marshal(customerDto)
				req, _ = http.NewRequest("POST", "/v1/customers", bytes.NewReader(jsonBytes))
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})

			// THEN
			It("should return an HTTP status code of 201-Created", func() {
				Expect(w.Code).To(Equal(http.StatusCreated))
			})
			It("should return the customer data in the body", func() {
				createdCustomer := &customer.CustomerDTO{}
				_ = json.Unmarshal(w.Body.Bytes(), &createdCustomer)

				Expect(createdCustomer.Id).ToNot(BeEmpty())
				Expect(createdCustomer.StreetAddress).To(Equal(customerDto.StreetAddress))
				Expect(createdCustomer.FullName).To(Equal(customerDto.FullName))
			})
		})

		// --
		Describe("Calling the creation endpoint with an invalid DTO should return bad request", func() {
			var (
				req *http.Request
				w   *httptest.ResponseRecorder
			)

			// GIVEN an invalid customer body
			BeforeEach(func() {
			})

			// WHEN the customer creation endpoint is called
			JustBeforeEach(func() {
				req, _ = http.NewRequest("POST", "/v1/customers", bytes.NewReader([]byte("Not JSON")))
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})

			// THEN
			It("should return an HTTP status code of 400-Bad Request", func() {
				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
			It("should return an error body", func() {
				errorDTO := common.ErrorDTO{}
				_ = json.Unmarshal(w.Body.Bytes(), &errorDTO)
				Expect(errorDTO.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errorDTO.Url).To(Equal("/v1/customers"))
				Expect(errorDTO.Message).To(Equal("invalid character 'N' looking for beginning of value"))
				Expect(errorDTO.Type).To(Equal("BadRequestError"))
			})
		})

	})
})
