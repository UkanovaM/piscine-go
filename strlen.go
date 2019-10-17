package piscine

func StrLen(str string) int {
	c := 0
	for _, p := range str {
		if p == p {
			c++
		}

	}

	return c
}
