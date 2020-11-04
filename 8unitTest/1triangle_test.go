package main

import (
	"testing"
)

// 测试函数名写法：Test被测函数名(t *testing.T)
func TestTriangle(t *testing.T) {

	// 测试数据
	tests := []struct{a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	// 遍历测试数据，运行测试函数，对每个测试数据进行检测
	for _, tt := range tests {
		if actual := calTriangle(tt.a, tt.b); actual != tt.c {
			// t 提供了汇报错误的一个接口
			t.Errorf("calTriangle(%d, %d); got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
