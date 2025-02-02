package ut_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/quexer/utee"
)

var _ = Describe("Ut", func() {
	It("", func() {
		pass := "ygc123456"
		hpw, err := utee.PasswdHash(pass)
		Ω(err).To(Succeed())
		Ω(hpw).ToNot(BeEmpty())

		Ω(utee.VerifyPasswd(pass, hpw)).To(BeTrue())
	})
})
