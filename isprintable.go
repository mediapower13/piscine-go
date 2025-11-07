package piscine

// IsPrintable returns true if s contains only printable ASCII characters (32 to 126),
// or is empty. Non-printable runes (like newlines, tabs, control chars) return false.
func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
