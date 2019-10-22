package piscine

import "github.com/01-edu/z01"

func AlphaCount(str string) int {

	counter := 0

	for letter := range str {
		z01.PrintRune(letter)
	}
	z01.PrintRune(counter)
	z01.PrintRune(str)
}
