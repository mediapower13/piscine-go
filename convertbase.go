package piscine

// ConvertBase converts the number represented by nbr from baseFrom to baseTo.
// baseFrom and baseTo are strings whose characters represent digits in order.
// Assumes valid bases and non-negative numbers.
func ConvertBase(nbr, baseFrom, baseTo string) string {
	// prepare rune slices for bases
	bf := []rune(baseFrom)
	bt := []rune(baseTo)
	if len(bf) == 0 || len(bt) == 0 {
		return ""
	}
	// map runes of baseFrom to their values
	val := make(map[rune]int)
	for i, r := range bf {
		val[r] = i
	}
	// parse nbr in baseFrom
	var n int64 = 0
	baseFromLen := int64(len(bf))
	for _, r := range nbr {
		v, ok := val[r]
		if !ok {
			return "" // invalid digit
		}
		n = n*baseFromLen + int64(v)
	}
	// convert to baseTo
	baseToLen := int64(len(bt))
	if n == 0 {
		return string(bt[0])
	}
	var digits []rune
	for n > 0 {
		rem := int(n % baseToLen)
		digits = append(digits, bt[rem])
		n = n / baseToLen
	}
	// reverse digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}
