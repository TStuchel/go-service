package testutils_test

import (
	. "github.com/TStuchel/go-service/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

// ------------------------------------------------ TEST SPECIFICATIONS ------------------------------------------------

func TestStructUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestStructUtils Suite")
}

// Ginkgo BDD tests
var _ = Describe("TestStructUtils", func() {

	type InnerStruct struct {
		InnerStringProp1 string
		InnerIntProp1    int
	}

	// Test structure (contains each known type)
	type SampleStruct struct {
		StringProp1 string
		StringProp2 string
		IntProp1    int
		IntProp2    int
		InnerProp   InnerStruct
		//StringArray []string
	}

	// Populating a structure
	// when given a pointer to that structure
	// should populate all fields with random data
	Describe("Populating a structure", func() {
		Context("when given a pointer to that structure", func() {
			var (
				testStruct SampleStruct
			)
			BeforeEach(func() {
				testStruct = SampleStruct{}
				PopulateTestData(&testStruct)
			})
			It("should populate all string fields of the structure with random strings", func() {
				Expect(testStruct.StringProp1).ToNot(BeEmpty())
				Expect(testStruct.StringProp2).ToNot(BeEmpty())
				Expect(testStruct.StringProp1).ToNot(Equal(testStruct.StringProp2))
			})
			It("should populate all int fields of the structure with random integers", func() {
				Expect(testStruct.IntProp1).ToNot(BeZero())
				Expect(testStruct.IntProp2).ToNot(BeZero())
				Expect(testStruct.IntProp1).ToNot(Equal(testStruct.IntProp2))
			})
			It("should (recursively) populate all nested structures with random values", func() {
				Expect(testStruct.InnerProp.InnerStringProp1).ToNot(BeEmpty())
				Expect(testStruct.InnerProp.InnerIntProp1).ToNot(BeZero())
			})
			AfterEach(func() {
				//b, _ := json.Marshal(testStruct)
				//fmt.Println(string(b))
			})
		})
	})

})
