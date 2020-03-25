package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("\n2-1 变量定义\n")
	defineVariables()
	fmt.Println(aa, xx, yy, zz)

	fmt.Printf("\n2-2 内建变量类型\n")
	euler()
	triangle()
	fmt.Printf("\n2-3 常量与枚举\n")
	consts()
	enums()

	fmt.Printf("\n2-4 条件语句\n")
	branch()
	fmt.Println(eval(1, 2, "+"))
	fmt.Println(grade(98))

	fmt.Printf("\n2-5 循环\n")
	forExample()
	printFile("a.txt")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Printf("\n2-6 函数\n")
	fmt.Println(evalAB(2, 3, "x"))
	if result, err := evalAB(2, 3, "x"); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(div(13, 3))
	q, r := divAB(14, 4)
	fmt.Println(q, r)
	// fmt.Println(apply(divAB, 3, 4))
	fmt.Println("pow(3, 4) is:", applyAB(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	fmt.Printf("\n2-7 指针\n")
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println("a, b after swap is:", a, b)

	a, b = 3, 4
	a, b = swapAB(a, b)
	fmt.Println("a, b after swapAB is:", a, b)

	fmt.Printf("\n2-8 atomic\n")
	test()
}
