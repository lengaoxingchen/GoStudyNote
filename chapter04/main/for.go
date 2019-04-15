package main

import "fmt"

/**
Go语言中的所有循环类型均可以使用for关键字来完成
for range(键值循环) --直接获取对象的索引和数据
适合遍历数组,切片\字符串\map及通道(channel).通过for range遍历的返回值有一定的规律
1.数组/切片/字符串返回索引和值
2.map返回键和值
3.通道只返回通道内的值
*/

func main() {
	//1.遍历数组/切片 --获取索引和元素
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d,value:%d\n", key, value)
	}

	//2.遍历字符串--获得字符

	var str = "hello你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	//3.遍历map--获取map的键和值
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	//4.遍历通道(channal)--接受通道数据
	//for range 可以遍历通道,但是通道在遍历时,只输出一个值,即管道内的类型对应的数据
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	//在遍历中选择希望获取的变量
	for _, value := range m {
		fmt.Println(value)
	}
}
