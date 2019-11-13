package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	for idx, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
		fmt.Println(idx, arg)
	}
	fmt.Println(s)
}
