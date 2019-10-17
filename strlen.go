package piscine

func StrLen(str string) int {
	c := 0
	a := []rune(str)
	for _, p := range str {
		c = p + 1
	}
	return c
}
