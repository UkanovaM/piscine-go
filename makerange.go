package piscine

func MakeRange(min, max int) []int {

	answer := make([]int, size)

	for i := min; i < max; i++ {
		answer[i] = i
	}
	return answer
}
