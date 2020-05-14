package http

import (
	"encoding/json"
	"errors"
	"github.com/TStuchel/go-service/common"
	"github.com/TStuchel/go-service/testutils"
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

	// --
	Context("HandleSuccess", func() {

		// --
		Describe("Handling success with no body content should return correct status", func() {
			var (
				w testutils.MockHttpResponseWriter
			)

			// GIVEN an HTTP response AND an HTTP status AND structure to return in the body
			BeforeEach(func() {
				w = testutils.MockHttpResponseWriter{}
			})

			// WHEN the response is handled
			JustBeforeEach(func() {
				HandleSuccess(&w, http.StatusAccepted, nil)
			})

			// THEN
			It("should set the HTTP status", func() {
				Expect(w.StatusCode).To(Equal(http.StatusAccepted))
			})
			It("should set the content header to JSON", func() {
				Expect(w.Header().Get("Content-Type")).To(Equal("application/json; charset=UTF-8"))
			})
			It("should return no body", func() {
				Expect(string(w.Bytes)).To(Equal(""))
			})

		})

		// --
		Describe("Handling success with body content should return correct status and the body as JSON", func() {
			var (
				w    testutils.MockHttpResponseWriter
				body interface{}
			)

			// GIVEN an HTTP response AND an HTTP status AND structure to return in the body
			BeforeEach(func() {
				w = testutils.MockHttpResponseWriter{}
				body = &struct {
					Data string
				}{Data: "value"}
			})

			// WHEN the response is handled
			JustBeforeEach(func() {
				HandleSuccess(&w, http.StatusAccepted, body)
			})

			// THEN
			It("should set the HTTP status", func() {
				Expect(w.StatusCode).To(Equal(http.StatusAccepted))
			})
			It("should set the content header to JSON", func() {
				Expect(w.Header().Get("Content-Type")).To(Equal("application/json; charset=UTF-8"))
			})
			It("should return the body as a JSON string", func() {
				Expect(string(w.Bytes)).To(Equal("{\"Data\":\"value\"}\n"))
			})

		})
	})

	// --
	Context("HandleNotFound", func() {

		// --
		Describe("Handling a not found error should return 404-Not Found", func() {
			var (
				w testutils.MockHttpResponseWriter
			)

			// GIVEN an HTTP response AND an HTTP status AND structure to return in the body
			BeforeEach(func() {
				w = testutils.MockHttpResponseWriter{}
			})

			// WHEN the response is handled
			JustBeforeEach(func() {
				HandleNotFound(&w)
			})

			// THEN
			It("should set the HTTP status", func() {
				Expect(w.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	// --
	Context("HandleBadRequest", func() {

		// --
		Describe("Handling an error should return 401-Unauthorized and contain the error data", func() {
			var (
				err      error
				w        testutils.MockHttpResponseWriter
				errorDTO *common.ErrorDTO
			)

			// GIVEN an (unauthorized) error
			BeforeEach(func() {
				err = errors.New("invalid data")
				w = testutils.MockHttpResponseWriter{}
			})

			// WHEN the error is handled
			JustBeforeEach(func() {
				HandleBadRequest(&w, "/endpoint", err)
				errorDTO = &common.ErrorDTO{}
				_ = json.Unmarshal(w.Bytes, errorDTO)
			})

			// THEN
			It("should set the HTTP to 400-Bad request", func() {
				Expect(w.StatusCode).To(Equal(http.StatusBadRequest))
			})
			It("should contain the expected error data", func() {
				Expect(errorDTO.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errorDTO.Url).To(Equal("/endpoint"))
				Expect(errorDTO.Message).To(Equal("invalid data"))
				Expect(errorDTO.Type).To(Equal("BadRequestError"))
			})
		})
	})

	// --
	Context("HandleUnauthorizedError", func() {

		// --
		Describe("Handling an unauthorized error should return 401-Unauthorized and contain the error data", func() {
			var (
				err      error
				w        testutils.MockHttpResponseWriter
				errorDTO *common.ErrorDTO
			)

			// GIVEN an (unauthorized) error
			BeforeEach(func() {
				err = errors.New("invalid credentials")
				w = testutils.MockHttpResponseWriter{}
			})

			// WHEN the error is handled
			JustBeforeEach(func() {
				HandleUnauthorizedError(&w, "/endpoint", err)
				errorDTO = &common.ErrorDTO{}
				_ = json.Unmarshal(w.Bytes, errorDTO)
			})

			// THEN
			It("should set the HTTP to 401-Unauthorized", func() {
				Expect(w.StatusCode).To(Equal(http.StatusUnauthorized))
			})
			It("should contain the expected error data", func() {
				Expect(errorDTO.StatusCode).To(Equal(http.StatusUnauthorized))
				Expect(errorDTO.Url).To(Equal("/endpoint"))
				Expect(errorDTO.Message).To(Equal("invalid credentials"))
				Expect(errorDTO.Type).To(Equal("UnauthorizedError"))
			})

		})
	})

})
