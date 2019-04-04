package main

import (
	"fmt"
	"sync"
)

/**
sync.Map是一种效率较高的并发安全
特性:
1.无须初始化,直接声明即可
2.sync.Map不能使用map的方式进行取值和设置等操作,而是使用sync.Map的方式进行调用.Store表示存储,Load表示获取,Delete表示删除
3.使用Range配合一个回调函数进行遍历操作,通过回调函数返回内部遍历出来的值.
Range参数中的回调函数的返回值功能是:
需要继续迭代遍历时,返回true;终止迭代遍历时,返回false
*/

func main() {
	var scene sync.Map
	//将键值对保存到sync.Map中
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	//根据键删除对应的键值对
	scene.Delete("london")

	//遍历所有sync.Map中的键值对
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate", key, value)
		return true
	})
}
