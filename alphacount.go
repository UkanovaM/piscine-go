package piscine

import "github.com/01-edu/z01"

func AlphaCount(x string) int {
	x := "Hello 78 World!    4455 /"

	counter := 0

	for letter := range x {

		counter++

		z01.AlphaCount("letter:%c\n", letter)
	}

	z01.AlphaCount("Counter value:%v\n", counter)
	z01.PrintRune(x)
}
