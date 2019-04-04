package main

import "fmt"

func main() {
	const (
		FlagNone = 1 <<iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("% d %d %d %d\n",FlagNone,FlagRed, FlagGreen,FlagBlue)
	fmt.Printf("%b %b %b %b\n",FlagNone,FlagRed,FlagGreen,FlagBlue )
}
