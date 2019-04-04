package main

import "fmt"

type NewInt int
type IntAlias = int

func main() {
	var a NewInt
	fmt.Printf("a type:%T\n", a)
	var a2 IntAlias
	fmt.Printf("a2 type:%T\n", a2)

}
