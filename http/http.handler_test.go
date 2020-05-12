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

	Context("HandleUnauthorizedError", func() {
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
			HandleUnauthorizedError(&w, "/v1/token", err)
			errorDTO = &common.ErrorDTO{}
			_ = json.Unmarshal(w.Bytes, errorDTO)
		})
		It("should contain the expected error data", func() {
			Expect(errorDTO.StatusCode).To(Equal(http.StatusUnauthorized))
			Expect(errorDTO.Url).To(Equal("/v1/token"))
			Expect(errorDTO.Message).To(Equal("invalid credentials"))
			Expect(errorDTO.Type).To(Equal("UnauthorizedException"))
		})

	})
})
