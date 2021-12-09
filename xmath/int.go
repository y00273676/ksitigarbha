package xmath

import (
	"github.com/shopspring/decimal"
)

// AbsInt64 Int64 绝对值
func AbsInt64(src int64) int64 {
	y := src >> 63
	return (src ^ y) - y
}

//GrowthRate. 支持到百分比后两位小数点
func GrowthRate(current, last int64, withPercentage bool) string {

	if last == 0 {
		return "-"
	}
	var val = (current - last) * 10000 / AbsInt64(last)

	var withoutPercentage = decimal.New(val, -2).StringFixed(2)
	if withPercentage {
		return withoutPercentage + "%"
	}
	return withoutPercentage
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
