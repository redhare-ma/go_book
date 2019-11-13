package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃ ", c)

}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	//	fmt.Printf("%g\n", boilingF-FreezingC) //编译错误，类型不匹配
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//	fmt.Println(c == f) //编译错误，类型不匹配
	fmt.Println(c == Celsius(f))

	c1 := FToC(212.0)
	fmt.Printf("%v\n", c1)
	fmt.Printf("%s\n", c1)
	fmt.Println(c1)
	fmt.Println(c1.String())
	fmt.Printf("%g\n", c1)
	fmt.Println(float64(c1))
}
