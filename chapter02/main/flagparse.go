package main

import (
	"flag"
	"fmt"
)

/**
使用指针变量获取命令行的输入信息\
Go语言的flag包中,定义的指令以指针类型返回.通过学习flag包,可以深入了解指针变量在设计上的方便之处
使用命令:go run flagparse.go  --mode=fast     打印结果为:fast

 */

 //通过flag.String,定义一个mode变量,这个变量的类型是*string.后面3个参数分别如下.
 //参数名称:在给应用输入参数时,使用这个名称
 //参数值的默认值:与flag所使用的函数创建变量类型对应,String对应字符串.Int对应整型.Bool对应布尔型
 //参数说明:使用-help时,会出现在说明中

var mode = flag.String("mode", "", "processMode")

func main() {
	//解析命令行参数,并将结果写入创建的指令变量mode中
	flag.Parse()
	//输出命令行参数,打印mode指针所指向的变量
	fmt.Println(*mode)
}
