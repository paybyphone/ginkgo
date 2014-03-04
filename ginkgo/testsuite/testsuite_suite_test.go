package testsuite_test

import (
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"

	"testing"
)

func TestTestsuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsuite Suite")
}
