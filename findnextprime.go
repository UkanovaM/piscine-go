package piscine

func FindNextPrime(x int) int {
	check := false
	for {
		if x <= 1 {
			check = false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				check = false
			}
		}
		check = true
		if check == true {
			return x
		}
		x++
	}
}
