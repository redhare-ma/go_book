package main

import "fmt"

func main() {
	fmt.Println(gcd(12, 134))
	fmt.Println(fib(9))
}

func gcd(x, y int) int {
	for y != 0 {
		fmt.Println(x, y)
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
