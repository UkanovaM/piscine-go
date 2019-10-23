package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	length := 0
	temp := ""
	for i := range os.Args {
		length = i
	}
	for i := 1; i <= length; i++ {
		for j := i + 1; j <= length; j++ {
			if Compare(os.Args[i], os.Args[j]) == 1 {
				temp = os.Args[i]
				os.Args[i] = os.Args[j]
				os.Args[j] = temp
			}
		}
		PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}
}
