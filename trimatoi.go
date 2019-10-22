package piscine

import "github.com/01-edu/z01"

func TrimAtoi(s string) int {

	a := 0

	ok := false

	for_, b := range s {
		if c >= '0' && c <= '9' {
			ok = true
			break
		}
	}
	return ok

	change := false

	for _, b := range s {
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
	if check(s) == true {
		for _, b := range s {
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