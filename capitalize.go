package piscine

func Capitalize(s string) string {
	str := []byte(s)
	cap := true
	for i, l := range str {
		if 96 < l && l < 123 {
			if cap {
				str[i] = str[i] - 32
				cap = false
			}
		} else if 64 < l && l < 91 {
			if !cap {
				str[i] = str[i] + 32
			} else {
				cap = false
			}
		} else if 47 < l && l < 58 {
			cap = false
		} else {
			cap = true
		}

	}
	return string(str)
}
