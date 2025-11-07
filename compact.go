package piscine

// front and reslicing to the new length.
func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}
	s := *ptr
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[j] = s[i]
			j++
		}
	}
	*ptr = s[:j]
	return j
}
