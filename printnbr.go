package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {

	a := 1
	if n < 0 {
		a = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		b := (n / 10) * t
		if b != 0 {
			PrintNbr('b')
		}
		c := (n % 10 * t) + '0'
		z01.PrintRune(rune('c'))
	} else {
		z01.PrintRune('0')
	}

}
