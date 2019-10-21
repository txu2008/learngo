package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// bool, string
// (u)int, (u)in8t, (u)int16, (u)int32, (u)int64, uintptr
// byte
// rune （字符型），==32位的char
// float32, float64,
// complex64, complex128 复数

func euler() {
	// complex
	c := 3 + 4i
	fmt.Println(c, cmplx.Abs(c))

	// euler 欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 类型转换是强制的
func triangle() {
	var a, b int = 3, 4
	var c int
	// c = math.Sqrt(float64(a*a + b*b))   // error
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(a, b, c)
}
