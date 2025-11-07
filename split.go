package piscine

// Split splits the string s by the separator sep and returns the slice of
// substrings. If sep is an empty string, nil is returned.
func Split(s, sep string) []string {
	if sep == "" {
		return nil
	}
	var res []string
	start := 0
	sepLen := len(sep)
	for i := 0; i <= len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			res = append(res, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}
	// append the remainder (including when sep was not found at all)
	res = append(res, s[start:])
	return res
}
