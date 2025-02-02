package types

import (
	"fmt"
	"strconv"
	"strings"

	"easyslip.cc/mic-project-layout/pkg/ut"
	"github.com/pkg/errors"
)

const (
	RawPercentUnit    = 10000
	NewRawPercentUnit = 100
)

// Percent 百分比，精确为小数点后最多两位
// 例
// 20%=2000
// 2%=200
// 1.2%=120
// 0.08%=8
type Percent int

func (p Percent) Fmt() string {
	s := fmt.Sprintf("%.2f", float64(p)/RawPercentUnit)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func (p Percent) ToText() string {
	return p.Fmt() + "%"
}

func (p Percent) ToFloat() float64 {
	return ut.Round(float64(p)/RawPercentUnit, 2)
}

func ParsePercent(s string) (Percent, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	// 没有引用 consts 包的RMBUnit, 保持本包少依赖
	return NewPercent(f), nil
}

func NewPercent(f float64) Percent {
	return Percent(ut.Round(f*NewRawPercentUnit, 0))
}
