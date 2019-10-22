package piscine

import "github.com/01-edu/z01"

func AlphaCount(str string) int {

	counter := 0

	for letter := range str {
		z01.AlphaCount(letter)
	}
	z01.AlphaCount(counter)
	z01.PrintRune(str)
}
