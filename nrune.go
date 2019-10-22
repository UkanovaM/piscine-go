package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	return sentence[n-1]
}
