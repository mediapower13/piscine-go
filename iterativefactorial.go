package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	maxInt := int(^uint(0) >> 1)
	res := 1
	for i := 1; i <= nb; i++ {
		if res > maxInt/i {
			return 0
		}
		res *= i
	}
	return res
}
