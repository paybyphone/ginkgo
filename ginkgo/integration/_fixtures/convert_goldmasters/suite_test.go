package convert_fixtures_test

import (
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"

	"testing"
)

func TestConvert_fixtures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert_fixtures Suite")
}
