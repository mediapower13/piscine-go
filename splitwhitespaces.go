package piscine

// SplitWhiteSpaces separates a string into words using spaces, tabs and newlines
// as separators and returns the slice of words. Consecutive separators are
// treated as a single separator. Leading and trailing separators are ignored.
func SplitWhiteSpaces(s string) []string {
	var res []string
	wordStart := -1
	for i, ch := range s {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if wordStart != -1 {
				res = append(res, s[wordStart:i])
				wordStart = -1
			}
		} else {
			if wordStart == -1 {
				wordStart = i
			}
		}
	}
	if wordStart != -1 {
		res = append(res, s[wordStart:])
	}
	return res
}
