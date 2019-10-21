package findnextprime

func check(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(x int) int {
	for {
		if check(x) {
			return x
		}
		x++
	}
}
