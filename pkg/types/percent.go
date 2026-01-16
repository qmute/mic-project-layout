package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/qmute/mic-project-layout/pkg/ut"
)

const (
	// PercentUnit 百分比单位，用于存储和展示转换
	// 存储值 = 百分比数值 * PercentUnit
	// 例如: 20% = 20 * 100 = 2000
	PercentUnit = 100
)

// Percent 百分比类型，精确到小数点后最多两位
// 内部存储方式: 百分比数值 * 100
//
// 存储示例:
//
//	20%   -> 2000
//	2%    -> 200
//	1.2%  -> 120
//	0.08% -> 8
type Percent int

// Fmt 格式化输出百分比文本，去除尾部无效的零
// 例如: 2000 -> "20", 120 -> "1.2", 8 -> "0.08"
func (p Percent) Fmt() string {
	s := fmt.Sprintf("%.2f", float64(p)/PercentUnit)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// ToText 返回带百分号的文本表示
// 例如: 2000 -> "20%", 120 -> "1.2%"
func (p Percent) ToText() string {
	return p.Fmt() + "%"
}

// ToValue 返回百分比原始数值（保留两位小数）
// 例如: 2000 -> 20.0, 120 -> 1.2
func (p Percent) ToValue() float64 {
	return ut.Round(float64(p)/PercentUnit, 2)
}

// ToDecimal 将百分比转换为小数形式（用于计算，保留四位小数）
// 例如: 2000(20%) -> 0.2, 120(1.2%) -> 0.012
func (p Percent) ToDecimal() float64 {
	return ut.Round(float64(p)/PercentUnit/PercentUnit, 4)
}

// Int64 返回 int64 类型的存储值
func (p Percent) Int64() int64 {
	return int64(p)
}

// ParsePercent 从字符串解析百分比
// 支持格式: "20", "1.2", "0.08"
func ParsePercent(s string) (Percent, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return NewPercent(f), nil
}

// NewPercent 从浮点数创建百分比
// 例如: NewPercent(50) -> 5000 (表示50%)
func NewPercent(f float64) Percent {
	return Percent(ut.Round(f*float64(PercentUnit), 0))
}
