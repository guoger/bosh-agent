package udevdevice_test

import (
	. "github.com/cloudfoundry/bosh-agent/internal/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry/bosh-agent/internal/github.com/onsi/gomega"

	"testing"
)

func TestUdevdevice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Udevdevice Suite")
}
