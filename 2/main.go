package main

import "fmt"

func main() {
	fmt.Print("basic 001")
	defineVariables()
	fmt.Println(aa, xx, yy, zz)

	euler()
	triangle()
	consts()
	enums()

	branch()
	fmt.Println(eval(1, 2, "+"))
	fmt.Println(grade(98))
}
