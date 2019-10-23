package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	len := 0
	caps := false
	start := 1
	tempInt := 0
	for i := range os.Args {
		len = i
	}
	if len > 0 {
		if Compare(os.Args[1], "--upper") == 0 {
			caps = true
		}
		for i := start; i <= len; i++ {
			tempInt = BasicAtoi(os.Args[i])
			if tempInt >= 1 && tempInt <= 26 {
				if caps {
					tempInt += 64
				} else {
					tempInt += 96
				}
				z01.PrintRune(rune(tempInt))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
