package main

// 函数可以返回多个值
// go:函数式编程
// 函数作为参数
// 没有默认参数和可变参数
// 可变参数列表， sum(1, 2, 3, 4, 5)

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数可以返回多个值
func evalAB(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// result = a / b
		q, _ := divAB(a, b)
		return q, nil
	default:
		// panic("unsupport operator:" + op)
		return 0, fmt.Errorf("unsupport operator:%s", op)
	}
}

// 13 / 3 = 4 .. 1
// 返回2 int
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 返回 q r
func divAB(a, b int) (q, r int) {
	// return a / b, a % b //或者
	q = a / b
	r = a % b
	return
}

// go:函数式编程
// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func applyAB(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 没有默认参数，可变参数
// 可变参数列表， sum(1, 2, 3, 4, 5)
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}
