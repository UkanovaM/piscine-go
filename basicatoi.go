package piscine

import (
	
)
func BasicAtoi(s string) int
   x := 0
   for _,c := range s {
	   cnt := 0
	   for i := '1'; i <= c; i++ {
		   cnt++
	   }
	   x = x*10 + cnt
   }
   return x 
}
