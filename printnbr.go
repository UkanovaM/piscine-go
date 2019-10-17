package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {

	k := 1
	if n < 0 {
		k = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		l := (n / 10) * k
		if l != 0 {
			PrintNbr('l')
		}
		m := (n % 10 * k) + '0'
		z01.PrintRune(rune('m'))
	} else {
		z01.PrintRune('0')
	}

}
