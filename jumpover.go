package piscine

func JumpOver(str string) string {
	var res []rune
	for i, r := range []rune(str) {
		if (i+1)%3 == 0 {
			res = append(res, r)
		}
	}
	return string(res) + "\n"
}
