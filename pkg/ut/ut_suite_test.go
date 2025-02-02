package ut_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ut Suite")
}
