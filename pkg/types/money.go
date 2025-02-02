package types

import (
	"fmt"
	"strings"

	"github.com/qmute/mic-project-layout/pkg/ut"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Money 钱, 单位为分
// 例:
// 1元 = 100
// 1角 = 10
// 1分 = 1
type Money int

// ToFloat 以元为单位，自动4舍5入至分位
func (p Money) ToFloat() float64 {
	return ut.Round(float64(p)/100, 2)
}

func (p Money) Val() int {
	return int(p)
}

func (p Money) Abs() Money {
	if p < 0 {
		return -p
	}
	return p
}

// FmtThousands 格化为千分号
func (p Money) FmtThousands() string {
	printer := message.NewPrinter(language.Chinese)
	s := printer.Sprintf("%f", p.ToFloat())
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// Fmt 格式化成元
// raw 是否保留右边0及小数点
func (p Money) Fmt(raw ...bool) string {
	s := fmt.Sprintf("%.2f", float64(p)/100)
	if len(raw) > 0 && raw[0] {
		return s
	}
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// NewMoney 把以元单位的浮点型转成Money， 作必要的4舍5入
func NewMoney(f float64) Money {
	return Money(ut.Round(f*100, 0))
}
