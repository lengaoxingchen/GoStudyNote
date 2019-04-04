package main

import "fmt"

/**
通过命令行分析变量逃逸
 */
//本函数测试入口参数和返回值情况
func dummy(b int) int {
	//声明一个c复制进入参数并返回
	var c int
	c = b
	return c
}

//空函数,测试没有任何参数函数的分析情况
func void() {

}

func main() {
	//声明a变量,测试main中变量的分析情况
	var a int
	//调用void()函数,没有返回值,测试void调用后的分析情况
	void()
	//打印a变量的值和dummy()函数返回,测试函数返回值没有变量接受时的分析情况
	fmt.Println(a, dummy(0))
}
