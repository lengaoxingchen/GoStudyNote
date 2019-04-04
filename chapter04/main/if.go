package main

/**
条件判断if的使用
格式:if 表达式1{分支1} else if 表达式2{分支2}else{分支3}
Go语言规定与if匹配的左括号"{"必须与if和表达式放在同一行,如果尝试将"{"放在其他位置,将会触发编译错误
同理.与else匹配的"{"也必须与else在同一行,else也必须与上一个if或else if的右边的大括号在一行.
*/
func main() {
	//一般写法
	var ten int = 11
	if ten > 10 {
		println(">10")
	} else {
		println("<=10")
	}

	//特殊写法 if还有一种特殊的写法,可以在if表达式之前添加一个执行语句,再根据变量值进行判断
	if err := Connect(); err != "" {
		println(err)
		return
	}
}

/**
Connect是一个带有返回值的函数,err:=Connect()是一个语句,执行Connect后,将错误保存到err变量中
err!=""才是if的条件表达式,当err不为空时,打印错误并返回.
这种写法可以将返回值与判断放在一行进行处理,而且返回值的作用范围被限制在if,else语句组合中
提示:在编程中,变量在其实现了变量的功能后,作用范围越小,所造成的问题可能性越小,每一个变量代表一个状态,有状态的地方,
状态就会被修改,函数的局部变量只会影响一个函数的执行.但全部变量可能会影响所有代码的执行状态,
因此限制变量的作用范围对代码的稳定性有很大的帮助
*/

func Connect() string {
	return "nel;"
}