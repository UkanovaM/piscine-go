package piscine

func MakeRange(min, max int) []int {
	var answer []int
	if min >= max {
		return nil
	}

	answer := make([]int, max-min)

	for i := min; i < max; i++ {
		answer[i] = i + min
	}
	return answer
}
