package types_test

import (
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"
	. "github.com/paybyphone/ginkgo/types"
)

var _ = Describe("GuidGenerator", func() {
	It("should generate a random guid", func() {
		a := GenerateRandomID()
		b := GenerateRandomID()
		Ω(a).ShouldNot(BeEmpty())
		Ω(b).ShouldNot(BeEmpty())
		Ω(a).ShouldNot(Equal(b))

		IDRegexp := "[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}"
		Ω(a).Should(MatchRegexp(IDRegexp))
		Ω(b).Should(MatchRegexp(IDRegexp))
	})
})
