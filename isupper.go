package piscine

func IsUpper(str string) bool {
	for _, x := range str {
		if x < 'A' || x > 'Z' {
			return false
		}
	}
	return true
}
