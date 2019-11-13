package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		i := tempconv.Inches(t)
		m := tempconv.Meter(t)
		k := tempconv.Kg(t)
		p := tempconv.Pounds(t)
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n", m, tempconv.MToI(m))
		fmt.Printf("%s = %s\n", i, tempconv.IToM(i))
		fmt.Printf("%s = %s\n", p, tempconv.PToK(p))
		fmt.Printf("%s = %s\n", k, tempconv.KToP(k))
	}

}
