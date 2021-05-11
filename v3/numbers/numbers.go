package numbers

func IsPrime(a int) bool {
	if a <= 1 {
		return false
	}
	for i := 2; i <= a/2; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func IsDivisible(a, b int) bool {
	if a%b == 0 {
		return true
	} else {
		return false
	}
}
