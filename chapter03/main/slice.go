package main

import (
	"fmt"
)

/**
切片的操作
从数组或切片生成新的切片拥有如下特性:
1.取出的元素为:结束位置-开始位置
2.取出元素不包含结束位置对应的索引,切片最后一个元素用slice[len(slice)]
3.当缺省开始位置时,表示从连续区域开头到结束位置
4.当缺省结束位置时,表示从开始位置到整个连续区域末尾
5.两者同时缺省时,与切片本身等效
6.两折同时为0时,等效于空切片,一般用于切片复位
7.根据索引位置取切片slice元素值时,取值范围是(0～len(slice-1)),超界会报运行时错误.生成切片时,结束位置可以填写len(slice)
但不会报错
 */
func main() {
	//语法: slice [开始位置:结束位置]
	//slice表示目标切片对象
	//开始位置对应目标切片对象的开始索引
	//结束位置对应目标切片对象的结束索引
	var a = []int{1, 2, 3}
	fmt.Println(a, a[1:2])
	//1.从指定范围中生成切片
	var highRiseBuilding [30] int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	//区间
	fmt.Println(highRiseBuilding[10:15])
	//中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	//开头到中间的所有元素
	fmt.Println(highRiseBuilding[:2])

	//2.表示原有的切片
	fmt.Println(a[:])

	//3.重置切片,清空拥有的元素
	fmt.Println(a[0:0])

	//4.声明切片的格式:var name []T
	//name表示切片类型的变量名
	//T表示切片类型对应的元素类型

	//声明字符串切片
	var strList []string
	//声明整型切片
	var numList []int
	//声明一个空切片
	var numListEmpty = []int{}
	//输出3个切片
	fmt.Print(strList, numList, numListEmpty)
	//输出三个切片的大小
	fmt.Print(len(strList), len(numList), len(numListEmpty))
	//切片判定空的结果
	fmt.Print(strList == nil)
	fmt.Print(numList == nil)
	fmt.Print(numListEmpty == nil)

	//5.使用make()函数构造切片
	
}
