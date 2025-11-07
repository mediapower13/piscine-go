package piscine

func UltimateDivMod(a *int, b *int) {
	if b == nil || a == nil {
		return
	}
	if *b == 0 {
		*a = 0
		*b = 0
		return
	}
	quot := *a / *b
	rem := *a % *b
	*a = quot
	*b = rem
}
