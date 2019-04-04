package main

import "fmt"

/**
功能:复制切片元素到另一个切片
使用格式:
copy(dest Slice,src Slice []T) int
src Slice为数据来源切片
dest Slice为复制的目标.目标切片必须分配过空间且足够承载复制的元素个数.来源和目标的类型一致,
copy的返回值表示实际发生复制的元素个数.
*/

func main() {
	//设置元素数量为1000
	const elementCount = 1000
	//与分配足够多的元素切片
	srcData := make([]int, elementCount)
	//将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	//引用切片数据
	refData := srcData
	//预分配足够多的元素切片
	copyData := make([]int, elementCount)
	//将数据复制到新的切片空间中
	copy(copyData, srcData)
	//修改原始数据的第一个元素
	srcData[0] = 999
	//打印引用切片的第一个元素
	fmt.Println(refData)
	//打印复制切片的第一个元素和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	//复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", copyData[i])
	}

	//从切片中删除元素
	seq := []string{"a", "b", "c", "d", "e"}
	//指定删除位置
	index := 2
	//查看删除位置之前的元素和之后的元素
	fmt.Println(seq[index], seq[index+1:])
	//将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}
