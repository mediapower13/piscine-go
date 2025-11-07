package piscine

// ToUpper returns s with ASCII lowercase letters converted to uppercase.
// Non-lowercase runes are left unchanged.
func ToUpper(s string) string {
	r := []rune(s)
	for i, c := range r {
		if c >= 'a' && c <= 'z' {
			r[i] = c - ('a' - 'A')
		}
	}
	return string(r)
}
