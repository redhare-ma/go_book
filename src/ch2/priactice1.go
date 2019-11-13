package main

import "fmt"

func main() {
	x := "hello!"
	fmt.Println("outer: ", x)
	for i := 0; i < len(x); i++ {
		y := x[i]
		fmt.Println("middle", y)
		if y != '!' {
			z := y + 'A' - 'a'
			fmt.Printf("inner %c\n", z)
		}
	}
}
