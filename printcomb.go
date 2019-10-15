package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		z01.PrintRune('i')
		z01.PrintRune('\n')
		for k := i + 1; k <= 9; k++ {
			z01.PrintRune('k')
			z01.PrintRune('\n')
			for l := k + 1; l <= 9; l++ {
				z01.PrintRune('l')
				z01.PrintRune('\n')
			}
		}
	}
}
