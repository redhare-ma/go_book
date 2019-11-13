package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	x := newInt()
	y := newInt1()
	fmt.Println(x, y)

	a := new(int)
	b := new(int)
	fmt.Println(a == b)

	c := new(string)
	fmt.Println(*c == "")

	d := new(map[int]int)
	fmt.Println(d)

	m := map[int]int{1: 10, 2: 2, 3: 3}
	fmt.Println(m[1])

}

func newInt() *int {
	return new(int)
}

func newInt1() *int {
	var dummy int
	return &dummy
}
