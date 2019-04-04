package main

import "fmt"

func main() {
	var a byte = 'a'
	//使用fmt.printf中的%T动词可以输出变量的实际类型,
	fmt.Printf("%d %T\n", a, a)

	var b rune = '你'
	fmt.Printf("%d %T\n", b, b)

	/*
	切片:是一个拥有相同类型元素的可变长度的序列.切片的声明方式如下:
	var name []T
	T代表切片元素类型,可以是整型,浮点型,布尔型,切片,map,函数等
	切片的元素使用[]进行访问,在方括号中提供切片的索引即可访问元素,索引的范围从0开始,且不超过切片的最大容量
	 */
	c := make([]int, 3)
	c[0] = 1
	c[1] = 2
	c[2] = 3

	str := "hello world"
	fmt.Println(str[6:])
}
