package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
字符串应用
查询总结:
1.strings.Index:正向搜索子字符串
2.strings.LastIndex:反向搜索子字符串
3.搜索的起始位置可以通过切片偏移计算
修改总结:
1.Go语言的字符串是不可变的
2.修改字符串时,可以将字符串转换为[]byte进行修改
3.[]byte和string可以通过强制类型转换互换
 */

func main() {
	//1.计算字符串的长度
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1)) //输出16 表示tip1的字符个数为16
	tip2 := "忍者"
	fmt.Println(len(tip2)) //输出6 表示tip2的字符个数为6, go语言字符串都是以utf8格式保存,每个中文占用3个字节
	// 然而根据习惯,"忍者"的字符个数应该是2.
	// 	fmt.Println(utf8.RuneCountInString())统计unicode字符数量
	fmt.Println(utf8.RuneCountInString(tip2))
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))

	//2.遍历字符串,获取每一个字符串元素
	//2.1 遍历每个ASCII字符  使用for的数值循环进行遍历,直接取每个字符串的下标获取ASCII字符
	theme := "狙击start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}
	//2.2按unicode字符遍历字符串 使用 for range 遍历
	for _, s := range theme {
		fmt.Printf("Unicode:%c %d\n", s, s)
	}

	//3.获取字符串的某一段字符
	// "," 中英文的逗号返回的结果不一致,请注意. tracer和substr的符号不一致,会报错:找不到
	tracer := "死神来了，死神bye bye"
	comma := strings.Index(tracer, "，")
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])

	//4.修改字符串 无法直接修改每一个字符元素,只能通过重新构建新的字符串并赋值给原来的字符串变量实现
	angel := "Heros never die"
	angelBytes := []byte(angel)
	for i := 5; i <= 10; i++ {
		angelBytes[i] = ' '
	}
	fmt.Println(string(angelBytes))

	//5.连接字符串
	hammer := "吃我一锤"
	sickle := "死吧"
	//声明字节缓冲
	var stringBuffer bytes.Buffer
	//把字符写入缓存
	stringBuffer.WriteString(hammer)
	stringBuffer.WriteString(sickle)
	//将缓存以字符串形式输出
	fmt.Println(stringBuffer.String())

	//6.格式化 使用格式化函数
	//fmt.Sprintf(格式化样式,参数列表)
	//格式化样式:字符串形式,格式化动词以%开头
	//参数列表:多个参数以逗号分隔,个数必须与格式化样式中的个数--对应,否则运行时会报错
	var progress = 2
	var target = 8
	title := fmt.Sprintf("已采集%d个药草,还需要%d个完成任务", progress, target)
	fmt.Println(title)
	pi := 3.14159
	//按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)
	//匿名结构体声明,并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v'%+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T'%T\n", profile)
	}
