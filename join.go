package piscine

func Join(s []string, sep string) string {

	result := ""

	first := true

	for _, x := range s {
		if first {
			result = x
			first = false
		} else {
			result = result + sep + x
		}

	}
	return result
}
