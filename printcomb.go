package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for k := i + 1; k <= '9'; k++ {
			for l := k + 1; l <= '9'; l++ {
				z01.PrintRune(i)
				z01.PrintRune(k)
				z01.PrintRune(l)
				if i != '7' || k != '8' || l != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
