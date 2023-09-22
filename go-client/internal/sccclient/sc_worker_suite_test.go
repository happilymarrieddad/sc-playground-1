package sccclient_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSccClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SccClient Suite")
}
