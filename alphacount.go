package piscine

import "github.com/01-edu/z01"

func AlphaCount(str string) int {
	str := "Hello 78 World!    4455 /"

	count := 0

	for len(str) > 0 {
		r, size := z01.AlphaCount(str)
		count++
		z01.Printf("%c %v\n", r, size)
		str = str[:len(str)-size]
	}
	z01.PrintRune("count:", count)
}
