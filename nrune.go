package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	ln := 0
	var r rune
	for _, a := range sentence {
		ln++
		if ln == n {
			return a
		}
	}
	return r
}
