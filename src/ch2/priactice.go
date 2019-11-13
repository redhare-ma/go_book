package main

import "fmt"

func f() {}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
	//	fmt.Println(h)
}
