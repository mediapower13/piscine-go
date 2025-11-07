package piscine

// Non-letter runes are left unchanged.
func Rot14(s string) string {
	r := []rune(s)
	for i, c := range r {
		if c >= 'a' && c <= 'z' {
			r[i] = 'a' + (c-'a'+14)%26
		} else if c >= 'A' && c <= 'Z' {
			r[i] = 'A' + (c-'A'+14)%26
		}
	}
	return string(r)
}
