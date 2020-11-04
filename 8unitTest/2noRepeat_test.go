package main

import (
	"testing"
)

// 测试最长不重复子串函数
// 可以借助左侧绿色按钮的coverage选项查看代码覆盖率
// 也可以在命令行中写入  go test -coverprofile=c.out生成c.out文件
// 用 go tool cover 查看c.out的各种使用方法
// 最简单一种是 go tool cover -html=c.out，在浏览器中打开html查看代码覆盖率情况
func TestSubstr(t *testing.T)  {

	// 测试用例
	tests := []struct{
		s string
		ans int
	} {
		// 常规例子
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// 特殊例子
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},

		// 中文支持
		{"吕图是个大西瓜", 7},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s;" +
				"expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

// Benchmark 用来做性能测试，测试这个函数执行某一个测试用例
// 也可以用命令行来执行：go test -bench .


func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13 ; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8
	b.ResetTimer()
	// b.N提供一个执行次数的参数；执行了1448400，每次828ns 828 ns/op
	for i := 0; i < b.N; i++ {	//
		actual := lengthOfNonRepeatingSubStr2(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d",
				actual, s, ans)
		}
	}
}

// 使用 go test -bench . -cpuprofile cpu.out 命令生成 cpu.out文件
// 使用 go tool pprof cpu.out 命令变成交互环境
// 使用 web可以从cpu.out看到程序执行时每个环节的执行时间




