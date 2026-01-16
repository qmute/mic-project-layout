package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/qmute/mic-project-layout/pkg/ut"
)

const (
	PercentUnit     = 10000 // 计算时的单位
	ShowPercentUnit = 100   // 展示时的单位
)

// Percent 百分比，精确为小数点后最多两位
// 例
// 20%=2000
// 2%=200
// 1.2%=120
// 0.08%=8
type Percent int

func (p Percent) Fmt() string {
	s := fmt.Sprintf("%.2f", float64(p)/ShowPercentUnit)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func (p Percent) ToText() string {
	return p.Fmt() + "%"
}

func (p Percent) ToFloat() float64 {
	return ut.Round(float64(p)/PercentUnit, 2)
}

func (p Percent) Int64() int64 {
	return int64(p)
}

func ParsePercent(s string) (Percent, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return NewPercent(f), nil
}

// NewPercent 构建Percent对象
// 例子： 50% = 5000
func NewPercent(f float64) Percent {
	return Percent(ut.Round(f*float64(ShowPercentUnit), 0))
}
