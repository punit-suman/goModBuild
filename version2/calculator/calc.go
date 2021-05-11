package calculator

func Add(a ...float64) float64 {
	sum := 0.
	for _, v := range a {
		sum += v
	}
	return sum
}

func Subtract(a float64, b float64) float64 {
	return a - b
}

func Multiply(a ...float64) float64 {
	prod := 1.
	for _, v := range a {
		prod *= v
	}
	return prod
}

func Divide(a float64, b float64) float64 {
	return a / b
}

func Power(a float64, b float64) float64 {
	prod := 1.
	for i := 1.; i <= b; i++ {
		prod *= a
	}
	return prod
}

func Factorial(a uint) uint {
	if a == 1 {
		return a
	} else {
		return a * Factorial(a-1)
	}
}
