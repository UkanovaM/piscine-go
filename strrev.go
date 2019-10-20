package piscine

func StrRev(s string) string {

	var reverse string
	for a := len(s) - 1; a >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
