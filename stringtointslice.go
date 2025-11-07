package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	rs := []rune(str)
	out := make([]int, 0, len(rs))
	for _, r := range rs {
		out = append(out, int(r))
	}
	return out
}
