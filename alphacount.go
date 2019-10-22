package piscine

import "github.com/01-edu/z01"

func AlphaCount(str string) int {

	counter := 0

	for letter := range str {
		z01.Printf(letter)
	}
	z01.Printf(counter)
	z01.PrintRune(str)
}
