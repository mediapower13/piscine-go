package piscine

// ToLower returns s with ASCII uppercase letters converted to lowercase.
// Non-uppercase runes are left unchanged.
func ToLower(s string) string {
	r := []rune(s)
	for i, c := range r {
		if c >= 'A' && c <= 'Z' {
			r[i] = c + ('a' - 'A')
		}
	}
	return string(r)
}
