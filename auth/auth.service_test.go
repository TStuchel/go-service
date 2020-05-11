package auth

import (
	"github.com/TStuchel/go-service/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestAuthService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Service Suite")
}

// Ginkgo BDD tests
var _ = Describe("Auth Service", func() {

	// --
	Context("Generate token", func() {
		service := NewAuthService()

		// --
		Describe("An invalid username and password should return an error", func() {
			var (
				username string
				password string
				token    string
				err      error
			)
			// GIVEN an invalid username and password
			BeforeEach(func() {
				username = testutils.RandomString(20)
				password = testutils.RandomString(20)
			})

			// WHEN a token is requested
			JustBeforeEach(func() {
				token, err = service.Login(username, password)
			})

			It("should return an empty token", func() {
				Expect(token).To(BeEmpty())
			})
			It("should return an error", func() {
				Expect(err.Error()).To(Equal("invalid credentials"))
			})
		})

		// --
		Describe("An valid username and password should return a JWT token", func() {
			var (
				username string
				password string
				token    string
				err      error
			)
			// GIVEN an invalid username and password
			BeforeEach(func() {
				username = "admin"
				password = "admin"
			})

			// WHEN a token is requested
			JustBeforeEach(func() {
				token, err = service.Login(username, password)
			})

			It("should return a JWT token", func() {
				Expect(token).ToNot(BeEmpty())
			})
			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})

	})

})
