package main

import "fmt"

const GoUsage = `Go is a tool for managing Go source code.
Usage:
	go command [arguments
	..]`

func main() {
	s := "left foot"
	t := s
	s += ", right foot"
	a := `\n\t...\r`
	fmt.Println("new s: ", s)
	fmt.Println("t: ", t)
	fmt.Println("a: ", a)
	fmt.Println("GoUsage: ", GoUsage)
}
