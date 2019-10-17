package piscine

func UltumateDivMod(a *int, b *int) {

	c := *a / *b
	*b = *a % *b
	*a = c
}
