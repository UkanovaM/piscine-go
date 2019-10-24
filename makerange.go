package piscine

func MakeRange(min, max int) []int {
	var answer []int
	if min >= max {
		return nil
	}

	answer := make([]int, max-min)

	for i := range answer
		answer[i] = i + min
	}
	return answer
}
