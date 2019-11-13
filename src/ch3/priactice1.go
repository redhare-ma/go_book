package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o \n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X \n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	s := "hello world!"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
}
