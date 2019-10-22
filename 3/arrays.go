package main

import "fmt"

// 数组是值类型
// [10]int 和 [20]int 是不同的类型
// 调用func f(arr [10]int)会拷贝数组，其他语言大多数是引用传递，会改变数组内容
// go语言中一般不直接使用数组， 而是使用切片

func printArray(arr [3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func defineArray() {
	var array1 [3]int
	array2 := [3]int{1, 2, 3}
	array3 := [...]int{7, 8, 9}
	var grid [2][3]bool

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(grid)

	printArray(array3)

	for i, v := range array3 {
		fmt.Println(i, v)
	}

}
