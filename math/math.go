package math

// Power calculate power
func Power(b, x int) (result int) {
	result = 1
	for i := 1; i <= x; i++ {
		result *= b
	}
	return result
}
