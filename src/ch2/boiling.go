package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %g℃\n", f, c)
	//print
	//boiling point = 212F or 100C
}
