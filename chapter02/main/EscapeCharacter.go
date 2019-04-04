package main

import (
	"fmt"
)

func main(){
	fmt.Println("str := \"c\\Go\\bin\\go.exe\"")

	//定义多行字符串,用数字1旁边反引号
	const str  = `第一行
	第二行
	第三行
	\r\n`
	fmt.Println(str)

}