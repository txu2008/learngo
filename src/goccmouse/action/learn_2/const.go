// 3
package main

import (
	"fmt"
	"math"
)

const q, w = 1, 2
const (
	a, b = 3, 4
)

// const 定义常量可以不指定类型
func consts() {
	const filename string = "a.txt"
	const a, b int = 3, 4
	const x, y = 2, 3
	var c, d int

	c = int(math.Sqrt(float64(a*a + b*b)))
	d = int(math.Sqrt(x*x + y*y))
	fmt.Println("const", c, d)
}

// 枚举
func enums() {
	const (
		a  = 1
		bb = 2
		c  = 3
	)
	const (
		x = iota
		_
		y
		z
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
