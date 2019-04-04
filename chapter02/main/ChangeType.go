package main

import (
	"fmt"
	"math"
)

/*
使用类型前置加括号的方式进行类型转换,一般格式如下:
T(表达式)
T代表转换的类型.表达式包括变量,复杂算子和函数返回值等
类型转换时,需要考虑两种类型的关系和范围,是否会发生数值截断等
浮点数在转换为整型时,会将小数部分去掉,只保留整数部分
整型截断在类型转换中发生的较为隐形,有些即为难追查的问题,很小一部分是由整型截断造成
 */
func main() {
	//输出个数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	//初始化一个32位整型值
	var a int32 = 1047483647
	//输出变变量的十六进制值和十进制值
	//%x 动词将数值以十六进制格式输出
	fmt.Printf("int32:0x%x %d\n", a, a)

	//将a变量数值转换为十六进制,发生数值截断
	b := int16(a)
	//输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	//将常量保存为float类型
	var c float32 = math.Pi

	//转换为int类型,浮点发生精度丢失
	fmt.Println(int(c))
}
