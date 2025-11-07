package piscine

// Index returns the index of the first instance of toFind in s, or -1 if toFind is not present.
// Behavior matches strings.Index: an empty toFind returns 0.
func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	ls := len(s)
	lf := len(toFind)
	if lf > ls {
		return -1
	}
	for i := 0; i <= ls-lf; i++ {
		j := 0
		for j < lf && s[i+j] == toFind[j] {
			j++
		}
		if j == lf {
			return i
		}
	}
	return -1
}
