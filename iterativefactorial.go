package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 0; i <= nb; i++ {
		result = result * i
	}
	return result
}
