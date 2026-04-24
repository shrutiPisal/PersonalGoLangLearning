package address_api_tests_test

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecondScenario", func() {
	Context("When ", func() {
		BeforeEach(func() {
			log.Println("We are inside before each")
		})
		It("addition should be as expected", func() {
			a := 2 + 3
			Expect(a).To(Equal(5))
		})

		It("Multiplication should be as expected", func() {
			m := 3 * 3
			Expect(m).To(Equal(9))
		})

		AfterEach(func() {
			log.Println("we are AfterEach")
		})
	})
})
