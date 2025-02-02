package types_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/qmute/mic-project-layout/pkg/types"
)

var _ = Describe("Month", func() {
	DescribeTable("", func(month string, want int) {
		days := types.Month(month).Days()
		Ω(days).To(Equal(want))
	},
		Entry("202401", "202401", 31),
		Entry("202402", "202402", 29),
		Entry("202406", "202406", 30),
		Entry("202411", "202411", 30),
		Entry("202412", "202412", 31),
		Entry("202302", "202302", 28),
	)

	It("strings.Repeat", func() {
		s := strings.Repeat("0", types.Month("202402").Days())
		Ω(s).To(Equal("00000000000000000000000000000"))
	})
})
