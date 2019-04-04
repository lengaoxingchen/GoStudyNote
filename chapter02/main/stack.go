package main

import "fmt"

/**
栈和堆的使用
 */

func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10

	return x
}

func main() {
	d := calc(3, 5)
	fmt.Println(d)
}
