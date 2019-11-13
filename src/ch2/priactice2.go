package main

import "fmt"

func main() {
	x := "hello!"
	fmt.Println(int('A'), int('a'))
	for _, x := range x {
		fmt.Print(int(x))
		fmt.Printf(" == %c\n", x)
		x := x + 'A' - 'a'
		fmt.Print("", int(x))
		fmt.Printf(" == %c\n", x)
	}
}
