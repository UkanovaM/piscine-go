package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	f := true
	for _, a := range arg {
		if !f {
			for _, j := range a {
				z01.PrintRune(j)
			}
			z01.PrintRune('\n')
		}
		f = false

	}
}
