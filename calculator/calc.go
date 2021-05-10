package calculator

func Add(a ...float64) float64 {
	sum := 0.
	for _, v := range a {
		sum += v
	}
	return sum
}
