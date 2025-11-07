package piscine

// TrimAtoi converts the digits found in s into an int.
// A leading '-' or '+' sign that appears before any digit determines the sign.
// If no digits are present, returns 0.
func TrimAtoi(s string) int {
	sign := 1
	signSeen := false
	digitsFound := false
	res := 0
	for _, r := range s {
		if r == '+' || r == '-' {
			if !digitsFound && !signSeen {
				if r == '-' {
					sign = -1
				}
				signSeen = true
			}
			continue
		}
		if r >= '0' && r <= '9' {
			digitsFound = true
			res = res*10 + int(r-'0')
		}
	}
	if !digitsFound {
		return 0
	}
	return sign * res
}
