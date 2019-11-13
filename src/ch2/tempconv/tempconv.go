package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvins float64
type Inches float64
type Meter float64
type Kg float64
type Pounds float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃ ", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉ ", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%gK", k)
}

func (i Inches) String() string {
	return fmt.Sprintf("%g英寸", i)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (k Kg) String() string {
	return fmt.Sprintf("%gKG", k)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%g磅", p)
}
