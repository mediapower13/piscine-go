package piscine

func DivMod(a int, b int, div *int, mod *int) {
	if b == 0 {
		// undefined, set both to zero to avoid panic
		*div = 0
		*mod = 0
		return
	}
	*div = a / b
	*mod = a % b
}
