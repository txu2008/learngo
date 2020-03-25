package sections

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// for 的条件里面不需要括号
// for 的条件里可以省略初始条件/结束条件/递增表达式
// example:

func forExample() {
	sum := 0
	for i := 1; i < 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

// 无 while关键字，实现while True： for {xxx} 或者 for aa {bb;}
// 省略初始条件/结束条件/递增表达式
func forever() {
	for {
		fmt.Println("abc")
	}
}

// 省略初始条件
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 省略结束条件
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}
