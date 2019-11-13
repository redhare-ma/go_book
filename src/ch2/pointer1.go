package main

import "fmt"

func main() {
	fmt.Println(f(), f())
	fmt.Println(f() == f())

	a := 1
	incr(&a)
	fmt.Println(a)
	fmt.Println(incr(&a))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
