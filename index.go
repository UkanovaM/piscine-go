package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	r := []rune(s)
	len := 0
	lenf := 0
	for i := range r {
		len = i
	}

	for i := range toFind {
		lenf = i
	}

	for i := 0; i < len-lenf; i++ {
		if s[i:i+lenf+1] == toFind {
			return i
		}
	}
	return -1