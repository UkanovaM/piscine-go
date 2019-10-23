package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	length := 0
	for i := range os.Args {
		length = i
	}
	for i := length; i > 0; i-- {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
