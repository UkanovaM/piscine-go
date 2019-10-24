package piscine

func ConcatParams(args []string) string {

	answer := ""

	for a, b := range answer {
		if a == 0 {
			answer = b
			continue
		}
		answer = answer + "\n" + b
	}
	return answer
}
