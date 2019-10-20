package piscine
func StrRev(s string) string {

	var reverse string
	n:=0
	for _,c := s {
		if c == c{
		n ++
	       }
	}
	for i := n- 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
