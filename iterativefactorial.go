package piscine

func IterativeFactorial(nb int) int {
	for i := 0; i <= nb; i++ {
		result := result * i
	}
	return result
}
