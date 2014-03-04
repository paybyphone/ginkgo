package integration_test

import (
	"fmt"
	. "github.com/onsi/gomega"
	. "github.com/paybyphone/ginkgo"
	"os/exec"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	installGinkgoCommand := exec.Command("go", "install", "github.com/paybyphone/ginkgo/ginkgo")
	err := installGinkgoCommand.Run()
	if err != nil {
		fmt.Printf("Failed to compile Ginkgo\n\t%s", err.Error())
	}

	RunSpecs(t, "Integration Suite")
}
