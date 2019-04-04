package main

import (
	"fmt"
	"sort"
)

/**
映射(map)--建立事物关联的容器
在业务和算法中需要使用任意忍耐性的关联关系时,就需要使用到映射.
Go语言提供的映射关系容器为map.map使用散列表(hash)实现
提示:大多数语言中映射关系容器使用两种算法:散列表和平衡树
散列表:查找复杂度为O(1),最坏的情况为O(n)
平衡树:查找复杂度始终是O(log n)
*/
func main() {
	//1.定义map的格式:
	//map[KeyType] ValueType   KeyType为键类型   ValueType是键对应的值类型
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	v := scene["route2"]
	fmt.Println(v)
	v, ok := scene["route"]
	if ok {
		fmt.Println(v)
	}
	//遍历map
	m := map[string]string{"W": "forward", "A": "left", "D": "right", "S": "backward"}
	var mList []string
	for n := range m {
		mList = append(mList, n)
		sort.Strings(mList)
	}
	fmt.Println(mList)

	//使用delete()函数从map中删除键值对
	//delete(map,键)
	delete(m, "W")
	fmt.Println(m)

	//清空map中的所有元素唯一方法就是重新make一个新的map
	//能够在并发环境中使用的map---sync.Map
	//Go语言中的map在并发情况下,只读是线程安全的.同时读写线程不安全
	o := make(map[int]int)

	//开启一段并发代码
	go func() {
		//不停地对map进行写入
		for {
			o[1] = 1
		}
	}()
	//开启一段并发代码
	go func() {
		for {
			_ = o[1]
		}
	}()
	//无线循环,让并发程序在后台执行
	for {
		fmt.Println("一")
	}

	//运行代码会报错:fatal error: concurrent map read and map write
	//并发的读写map.也就是说使用了两个并发函数不断地对map进行了读和写而发生了竞态问题.
	// map内部会对这种并发操作进行检查并提前发现
	//需要并发读写时,一般的做法是加锁,但这样性能并不高.
}
