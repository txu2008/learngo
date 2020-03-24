// 1
package main

import "fmt"

// 数组占用的内存是连续分配的，数组的每个元素类型相同，又是连续分配，因此以固定速度索引数组中的任意数据，速度非常快
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
	// 声明一个有5 个元素的数组
	// 用具体值初始化索引为1 和2 的元素
	// 其余元素保持零值
	array4 := [5]int{1: 10, 2: 20}
	array5 := [5]*int{1: new(int), 2: new(int)}
	*array5[1] = 10
	*array5[2] = 20

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5, *array5[1], *array5[2], &array5[0], &array5[3])
	fmt.Println(grid)

	printArray(array3)

	for i, v := range array3 {
		fmt.Println(i, v)
	}

}
