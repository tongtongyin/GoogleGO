package main

import (
	"fmt"
	"math"
)

// 打印直角三角形斜边值
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calTriangle(a, b))
}

// 输入两个直角边，返回直角三角形斜边值
func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func main() {
	triangle()
}
