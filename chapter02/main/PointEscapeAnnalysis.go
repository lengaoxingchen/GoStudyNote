package main

import "fmt"

/**
使用结构体做数据,了解在对上分配的情况
 */
type Data struct {

}

func dummp() *Data{
	//实例化c为Data类型
	var c Data

	//返回局部变量地址
	return &c
}

func main(){
	fmt.Println(dummp())
}