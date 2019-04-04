package main

import "fmt"

/*
指针类型:
类型指针,允许对这个指针类型的数据进行修改.传递数据使用指针,而无须拷贝数据.类型指针不能进行偏移和运算
切片,由指向起始元素的原始指针,元素数量和容量组成
每个变量在运行时都拥有一个地址.这个地址代表变量在内存中的位置
Go语言中使用"&"操作符放在变量前面对变量进行"取地址"操作
格式如下:
ptr := &v   //v的类型为T
v代表被取地址的变量,被取地址的v使用ptr进行接受,ptr的类型就为"*T",称作T的指针类型."*代表指针".

取值:
在对普通变量使用"&"操作符获取这个变量的指针后,可以对指针使用"*" 操作,也就是指针取值


创建指针的另一种方法---new()函数
格式:new(类型)
 */

func main() {

	/**
	获取指针的地址
	 */
	var cat int = 1

	var str string = "banana"

	//%p输出变量取地址后的指针值,指针带有"0x"的十六进制前缀
	fmt.Printf("%p %p\n", &cat, &str)

	/**
	获取指针的值
	 */
	var house = "Malibu Point 10880,90265"

	//对字符串取地址,ptr类型为* string
	ptr := &house

	//打印ptr的类型
	fmt.Printf(" ptr type:%T\n", ptr)

	//打印ptr的指针地址
	fmt.Printf("address:%p\n", ptr)

	//对指针进行取值操作
	value := * ptr

	//取值后的类型
	fmt.Printf("value type: %T\n", value)

	//指针去之后就是指向变量的值
	fmt.Printf("value:%s\n", value)

	x, y := 1, 2
	fmt.Println(x, y)
	fmt.Println(&x, &y)

	swap(&x, &y)
	fmt.Println(x, y)
	fmt.Println(&x, &y)

	m, n := 3, 4
	fmt.Println(&m, &n)
	swap1(&m,&n)
	fmt.Println(&m, &n)

	//用new()函数创建指针
	leg := new(string)
	*leg = "ninja"
	fmt.Println(*leg)

}

/**
交换指针的值
 */
func swap(a, b *int) {

	//取a指针的值,赋值给临时变量t
	t := *a

	//取b指针的值,赋值给a指针指向的变量
	*a = *b
	//将a指针的值赋给b指针指向的变量
	*b = t

}
/**
swap1函数中交换操作的是指针值
结果表明:指针值是不可交换的.
swap1函数交换的是a和b的地址,在交换完毕后,a和b的变量值确实被交换.但是a和b关联的两个变量并没有实际关联.
这就像写有两座房子的卡片放在桌上一字摊开,交换两座房子的卡片后并不会对两座房子有任何影响.
 */
func swap1(a, b *int) {
	b, a = a, b

}
