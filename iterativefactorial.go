package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb >= 0 && nb <= 20 {
		for i := 0; i <= nb; i++ {
			result = result * i
		}
	}
	return result
}
