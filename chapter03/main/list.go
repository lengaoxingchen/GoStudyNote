package main

import (
	"container/list"
	"fmt"
)

/**
列表(list) --可以快速增删的非连续空间的容器
列表由多个节点组成,节点之间通过一些变量记录彼此之间的关系
列表的实现:单链表和双链表
Go语言中,将列表使用container/list包来实现,内部实现原理是双链表.列表能够高效地

列表与切片和map不同的是,列表并没有具体元素类型的显示.因此,列表元素可以是任意类型.这既带来便利,也会引来一些问题
给一个列表放入了非期望类型的值,在取出值后,将interface{}转换为期望类型时将会发生宕机.

双链表支持从队列前方或者后方插入元素,分别对应的方法是PushFront和PushBack.
这两个方法都会返回一个*list.Element结构.
如果在以后的使用中需要删除插入的元素,则只能通过*list.Element配合Remove()方法进行删除,这种方法可以让删除更加效率化,也是双链表的热性之一

*/

func main() {
	//1.初始化列表
	//方式一:通过container/list包的New方法初始化list
	li := list.New()
	fmt.Printf("%T\n", li)
	//方式二:通过声明初始化list
	var lis list.List
	fmt.Printf("%T\n", lis)

	//2.在列表中插入元素
	l := list.New()
	l.PushBack("canon")
	l.PushFront(67)

	//3.从列表中删除元素:
	// 列表的插入函数的返回值会提供一个*list.Element结构,这个结构记录着元素的值,以及和其他节点之间的关系等信息
	//从列表中删除元素时,需要用到这个结构进行快速删除

	//尾部添加后保存元素句柄
	element := l.PushBack("fist")
	//在fist之后添加high
	l.InsertAfter("high", element)
	//在fist之间添加noon
	l.InsertBefore("noon", element)
	//使用
	l.Remove(element)

	//4.遍历列表--访问列表的每一个元素
	//遍历双链表需要配合Front()函数获取头元素,遍历时只要元素不为空就可以继续进行.每一次遍历调用元素的Next
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
