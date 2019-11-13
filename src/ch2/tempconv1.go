package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.KToC(328.15))
}
