// 2
package main

import "fmt"

func defineSlice() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[2:6]

	fmt.Println(a)
	fmt.Println(b)
}
