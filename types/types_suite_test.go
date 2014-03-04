package types_test

import (
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"

	"testing"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Types Suite")
}
