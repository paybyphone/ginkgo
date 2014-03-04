package nested_test

import (
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"

	"testing"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Suite")
}
