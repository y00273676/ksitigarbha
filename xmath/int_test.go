//author: zhengchao.deng zhengchao.deng@meican.com
//date: 2019/5/17
package xmath_test

import (
	"ksitigarbha/xmath"
	"testing"
)

func TestAbsInt(t *testing.T) {
	var src = []int64{1, 3, 4, 123414141, -342, 234, -2342, 0, -9999999999999999, 8888888888888}
	var expects = []int64{1, 3, 4, 123414141, 342, 234, 2342, 0, 9999999999999999, 8888888888888}
	for i := range src {
		var result = xmath.AbsInt64(src[i])
		if result != expects[i] {
			t.Fatalf("src :%v,abs expects to be :%v,but result is :%v", src[i], expects[i], result)
		}
	}
}

func TestGrowthRate(t *testing.T) {
	var current = []int64{606, 101, 100}
	var last = []int64{5, 5, 0}
	var expect1 = []string{"12020.00", "1920.00", "-"}
	var expect2 = []string{"12020.00%", "1920.00%", "-"}
	var result1, result2 string
	for i := 0; i < len(current); i++ {
		result1 = xmath.GrowthRate(current[i], last[i], false)
		result2 = xmath.GrowthRate(current[i], last[i], true)
		if result1 != expect1[i] {
			t.Fatalf("current: %v, last: %v,withPercentage: false,expect growthRate: %v,but result is:%v", current[i], last[i], expect1[i], result1)
		}
		if result2 != expect2[i] {
			t.Fatalf("current: %v, last: %v,withPercentage: true,expect growthRate: %v,but result is:%v", current[i], last[i], expect2[i], result2)
		}
	}
}
