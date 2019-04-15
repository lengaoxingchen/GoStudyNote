package main

import "fmt"

/**
跳转执行代码标签(goto):goto语句通过标签进行代码间的无条件跳转.goto语句可以在快速跳出循环,避免重复退出上有一定的帮助.
go语言中使用goto语句能简化一些代码的实现过程
*/
func main() {
	notUseGoto()
	useGoto()
	checkError()
}
func notUseGoto() {
	//1.使用goto退出多层循环.下面的代码在满足条件时,需要连续退出两层循环.传统实现方式如下:
	var breakAgain bool
	//外循环
	for x := 0; x < 10; x++ {
		//内循环
		for y := 0; y < 10; y++ {
			//满足某个条件时,退出循环
			if y == 2 {
				//设置退出标记
				breakAgain = true
				//退出本次循环
				break
			}

		}
		//根据标记,还需要退出一次循环
		if breakAgain {
			break
		}
	}
	fmt.Println("done")
}

func useGoto() {
	//使用goto语句优化后的代码
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				//跳转到标签
				goto breakHere
			}
		}
	}
	//手动返回,避免执行进入标签
	return
	//便签
breakHere:
	fmt.Println("done")
}
