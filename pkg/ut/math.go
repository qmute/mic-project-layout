package ut

import (
	"math"
)

// Round 四舍五入方法，float保留n位小数
func Round(f float64, n int) float64 {
	if f >= 0 {
		n10 := math.Pow10(n)
		return math.Trunc((f+0.5/n10)*n10) / n10
	}

	// 负数转成正数四舍五入，再转回负数，结果已与产品确认
	f = f * -1
	n10 := math.Pow10(n)
	f = math.Trunc((f+0.5/n10)*n10) / n10

	return f * -1
}
