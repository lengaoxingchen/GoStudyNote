package main

import "fmt"

/**
分支选择(switch)--拥有多条分支的判断
GO语言中的switch,不仅可以基于常量进行判断,还可以基于表达式进行判断
Go语言中switch的每一个case与case之间是独立的代码块,不需要通过break语句跳出当前case代码块以避免执行到下一行
*/
func main() {
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
	//1.一分支多值  当出现多个case要放在一起的时候,不同的case表达式使用逗号分隔
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}
	//2.分支表达式 case后不仅只是常量,还可以和if一样添加表达式
	//这种情况switch后面不再跟判断变量,连判断的目标都没有了
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	default:
		fmt.Println("0")
	}
	//3.跨越case的fallthrough--兼容C语言的case设计
	//Go语言的case是一个独立的代码块,执行完毕后不会像C语言那样紧接着下一个case执行.为了兼容一些移植代码,
	// 依然加入了fallthrough关键字来实现.新编代码不建议是用哪个fallthrough
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}
