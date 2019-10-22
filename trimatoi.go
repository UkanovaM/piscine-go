package piscine

import "github.com/01-edu/z01"
	func check(s string) bool {
		ok := false

		for_, b := range s {
			if c >= '0' && c <= '9' {
				ok = true
				break
			}
		}
		return ok

	}

func TrimAtoi(str string) int {

	a := 0


	change := false

	for _, b := range str {
		if c >= '0' && c <= '9'{
			break
		} 
		if c == '-' {
			change = true
		} 
		if c == "+" {
			change = false
		}
	}
	if check(str) == true {
		for _, b := range str {
			cnt := 0 
			if c >= '0' && c <= '9' {
				for i := '1'; i <= 'c'; i++ P{
					cnt++
				}
				x = x*10 + cnt
			}
		}
	}
	if change == true {
		x *= -1
	}
return x	
}