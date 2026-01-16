package types_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/qmute/mic-project-layout/pkg/types"
)

var _ = Describe("Percent", func() {

	Describe("Fmt()", func() {
		It("整数百分比返回无小数点格式", func() {
			Expect(types.NewPercent(20).Fmt()).To(Equal("20"))
		})

		It("一位小数返回一位小数格式", func() {
			Expect(types.NewPercent(1.2).Fmt()).To(Equal("1.2"))
		})

		It("两位小数返回两位小数格式", func() {
			Expect(types.NewPercent(0.08).Fmt()).To(Equal("0.08"))
		})

		It("去除尾部零", func() {
			Expect(types.NewPercent(1.5).Fmt()).To(Equal("1.5"))
			Expect(types.NewPercent(2.5).Fmt()).To(Equal("2.5"))
		})
	})

	Describe("ToText()", func() {
		It("整数百分比添加%符号", func() {
			Expect(types.NewPercent(20).ToText()).To(Equal("20%"))
		})

		It("小数百分比添加%符号", func() {
			Expect(types.NewPercent(1.2).ToText()).To(Equal("1.2%"))
		})

		It("两位小数百分比添加%符号", func() {
			Expect(types.NewPercent(0.08).ToText()).To(Equal("0.08%"))
		})
	})

	Describe("ToValue()", func() {
		It("整数百分比返回原始值", func() {
			Expect(types.NewPercent(20).ToValue()).To(Equal(20.0))
		})

		It("小数百分比返回原始值", func() {
			Expect(types.NewPercent(1.2).ToValue()).To(Equal(1.2))
		})

		It("两位小数百分比返回原始值", func() {
			Expect(types.NewPercent(0.08).ToValue()).To(Equal(0.08))
		})
	})

	Describe("ToDecimal()", func() {
		It("整数百分比转换为小数", func() {
			// 20% = 2000, 2000/100/100 = 0.2
			Expect(types.NewPercent(20).ToDecimal()).To(Equal(0.2))
		})

		It("小数百分比转换为小数", func() {
			// 1.2% = 120, 120/100/100 = 0.012
			Expect(types.NewPercent(1.2).ToDecimal()).To(Equal(0.012))
		})

		It("两位小数百分比转换为小数", func() {
			// 0.08% = 8, 8/100/100 = 0.0008
			Expect(types.NewPercent(0.08).ToDecimal()).To(Equal(0.0008))
		})

		It("较大值百分比转换正确", func() {
			// 50% = 5000, 5000/100/100 = 0.5
			Expect(types.NewPercent(50).ToDecimal()).To(Equal(0.5))
		})

		It("100百分比转换正确", func() {
			// 100% = 10000, 10000/100/100 = 1
			Expect(types.NewPercent(100).ToDecimal()).To(Equal(1.0))
		})
	})

	Describe("Int64()", func() {
		It("返回int64类型", func() {
			p := types.NewPercent(15.5)
			Expect(p.Int64()).To(Equal(int64(1550)))
		})
	})

	Describe("NewPercent()", func() {
		It("整数百分比正确转换", func() {
			p := types.NewPercent(20)
			Expect(p.Int64()).To(Equal(int64(2000)))
		})

		It("小数百分比正确转换", func() {
			p := types.NewPercent(1.2)
			Expect(p.Int64()).To(Equal(int64(120)))
		})

		It("两位小数百分比正确转换", func() {
			p := types.NewPercent(0.08)
			Expect(p.Int64()).To(Equal(int64(8)))
		})

		It("浮点数四舍五入", func() {
			p := types.NewPercent(1.235)
			Expect(p.Int64()).To(Equal(int64(124))) // 四舍五入到整数
		})
	})

	Describe("ParsePercent()", func() {
		It("解析整数百分比字符串", func() {
			p, err := types.ParsePercent("20")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Int64()).To(Equal(int64(2000)))
		})

		It("解析小数百分比字符串", func() {
			p, err := types.ParsePercent("1.2")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Int64()).To(Equal(int64(120)))
		})

		It("解析两位小数百分比字符串", func() {
			p, err := types.ParsePercent("0.08")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Int64()).To(Equal(int64(8)))
		})
	})

	Describe("类型零值", func() {
		It("零值格式化为0", func() {
			var p types.Percent
			Expect(p.Fmt()).To(Equal("0"))
		})

		It("零值ToText为0%", func() {
			var p types.Percent
			Expect(p.ToText()).To(Equal("0%"))
		})
	})
})
