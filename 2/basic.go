package main

import (
	"fmt"
)

var aa int = 01

// bb :=33   // error, ":=" 定义变量只能在函数内部使用
var (
	xx int    = 1
	yy string = "2"
	zz        = 3
)

func defineVariables() {
	var a, b, c int = 1, 2, 3
	var s string
	var s2 string = "ssss222"
	var x, y, z = 1, "yy", true //编译器可推测变量类型
	q, w, e := true, 1, "ww"

	fmt.Println(a, b, c)
	fmt.Printf("%d, %d, %d, %q, %s", a, b, c, s, s2)
	// fmt.Print(a, b, c, s, s2)
	fmt.Println(x, y, z)
	fmt.Println(q, w, e)
}
