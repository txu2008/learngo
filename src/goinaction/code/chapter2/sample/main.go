package main

import (
	"log"
	"os"

	_ "goinaction/code/chapter2/sample/matchers" // 下划线允许对包做初始化操作（调用init 函数），但是并不使用包里的标识符
	"goinaction/code/chapter2/sample/search"
)

// init is called prior to main.  init 在main 之前调用
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
