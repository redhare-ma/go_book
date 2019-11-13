package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println("x:", x)
	fmt.Println("&x: ", p)
	fmt.Println("*p: ", *p)
	*p = 2
	fmt.Println("x:", x)
	fmt.Println("p:", p != nil)

	var z, y int
	fmt.Println(&z == &z, &z == &y, &z != nil)
}
