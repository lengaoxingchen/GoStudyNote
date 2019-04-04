package main

import "fmt"

/**
数组的使用
 */
func main() {
	//1.数组的声明.语法 var 数组变量名[元素数量] T
	//数组变量名:数组声明以及使用时的变量名
	//元素数量:数组的元素数量.可以是一个表达式,但是最终通过编译器计算的结果必须是整型数值.
	// 也就是说,元素数量不能含有运行时才能确认大小的数值
	//T可以是任意基本类型,包括T为数组本身.但类型为数组本身时,可以实现多维数组

	//方式一:
	//var team [3]string
	//team[0] = "hammer"
	//team[1] = "soldier"
	//team[2] = "mum"
	//方式二:数组大小可以不写,也可以写成"...",编译时编译器会根据元素个数确定数组大小
	//var team = [3]string{"hammer", "soldier", "mum"}
	//var team = []string{"hammer", "soldier", "mum"}
	var team = [...]string{"hammer", "soldier", "mum"}
	fmt.Println(team)

	//2.遍历数组
	for k, v := range team {
		fmt.Println(k, v)
	}
}
