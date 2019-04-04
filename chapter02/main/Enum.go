package main

import "fmt"

func main() {

	type Weapon int

	const (
		Arrow Weapon = iota
		//开始生成枚举值,默认为0
		Shrunken
		SniperRifle
		Rifle
		Blower
	)
	//输出所有的枚举值
	fmt.Println(Arrow, Shrunken, SniperRifle, Rifle, Blower)
	//使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)


}
