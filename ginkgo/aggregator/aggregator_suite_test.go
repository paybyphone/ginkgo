package aggregator_test

import (
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"

	"testing"
)

func TestGinkgoAggregator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Aggregator Suite")
}
