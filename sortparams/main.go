package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	ln := 0
	for c := range arg {
		ln = c
	}
	arg = arg[1 : ln+1]

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for _, c := range arg {
		for _, j := range c {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
