package main

func AppendRange(min, max int) []int {

	var answer []int

	for i := min; i < max; i++ {
		answer = append(answer, i+1)
	}
	return answer
}
