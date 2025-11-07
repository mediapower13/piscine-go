package piscine

// IsAlpha returns true if s is empty or contains only alphanumeric ASCII characters (A-Z, a-z, 0-9).
func IsAlpha(s string) bool {
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			continue
		}
		return false
	}
	return true
}
