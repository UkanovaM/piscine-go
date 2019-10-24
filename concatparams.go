package piscine

func ConcatParams(args []string) string {

	answer := ""

	for a, i := range args {
		if a == 0 {
			answer = i
			continue
		}
		answer = answer + "\n" + i
	}
	return answer
}
