package piscine

func TrimAtoi(s string) int {
	num := []rune{}
	len := 0
	positive := false
	negative := false
	for _, l := range s {
		if l == '-' && !positive {
			negative = true
		}
		if l > 47 && l < 58 {
			num = append(num, l)
			len++
			if !negative {
				positive = true
			}
		}
	}
	d := 1
	ans := 0
	for i := len - 1; i > -1; i-- {
		ans = ans + (int(num[i])-48)*d
		d = d * 10
	}
	if negative {
		ans = ans * (-1)
	}
	return ans
}
