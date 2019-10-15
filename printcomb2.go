package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for k := i + 1; k <= '9'; k++ {
			z01.PrintRune(i)
			z01.PrintRune(k)
			if i != '7' || k != '8' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
