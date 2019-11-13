package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}

func MToI(m Meter) Inches {
	return Inches(m * 39.3700787)
}

func IToM(i Inches) Meter {
	return Meter(i * 0.0254)
}

func PToK(p Pounds) Kg {
	return Kg(p * 0.45359237)
}

func KToP(k Kg) Pounds {
	return Pounds(k * 2.20462262)
}
