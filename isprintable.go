package piscine

func IsPrintable(str string) bool {
	for _, x := range str {
		if x < 32 || x > 127 {
			return false
		}
	}
	return true
}
