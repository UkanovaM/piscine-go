package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, x := range arg[0] {
		z01.PrintRune(x)
	}
	z01.PrintRune(10)
}
