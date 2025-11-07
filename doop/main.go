package main

import "os"

const (
	maxInt64 = uint64(9223372036854775807)
	minInt64 = -9223372036854775808
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, ok := parseInt64(aStr)
	if !ok {
		return
	}
	b, ok := parseInt64(bStr)
	if !ok {
		return
	}

	switch op {
	case "+":
		if addWillOverflow(a, b) {
			return
		}
		writeInt64(a + b)
	case "-":
		if addWillOverflow(a, -b) {
			return
		}
		writeInt64(a - b)
	case "*":
		if mulWillOverflow(a, b) {
			return
		}
		writeInt64(a * b)
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		writeInt64(a / b)
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		writeInt64(a % b)
	default:
		return
	}
}

func parseInt64(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	i := 0
	neg := false
	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		i = 1
		neg = true
	}
	if i >= len(s) {
		return 0, false
	}
	var u uint64 = 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := uint64(c - '0')
		limit := maxInt64
		if neg {
			limit = maxInt64 + 1
		}
		if u > limit/10 || (u == limit/10 && d > limit%10) {
			return 0, false
		}
		u = u*10 + d
	}
	if neg {
		if u == maxInt64+1 {
			return minInt64, true
		}
		return -int64(u), true
	}
	if u > maxInt64 {
		return 0, false
	}
	return int64(u), true
}

func addWillOverflow(a, b int64) bool {
	if b > 0 {
		return a > int64(maxInt64)-b
	}
	if b < 0 {
		return a < minInt64-b
	}
	return false
}

func absUint64(x int64) uint64 {
	if x < 0 {
		if x == minInt64 {
			return maxInt64 + 1
		}
		return uint64(-x)
	}
	return uint64(x)
}

func mulWillOverflow(a, b int64) bool {
	if a == 0 || b == 0 {
		return false
	}
	ua := absUint64(a)
	ub := absUint64(b)
	positive := (a > 0) == (b > 0)
	if positive {
		return ua > maxInt64/ub
	}
	return ua > (maxInt64+1)/ub
}

func writeInt64(n int64) {
	os.Stdout.WriteString(itoa(n))
	os.Stdout.WriteString("\n")
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	if n == minInt64 {
		return "-9223372036854775808"
	}
	neg := n < 0
	var u uint64
	if neg {
		u = uint64(-n)
	} else {
		u = uint64(n)
	}
	var buf [20]byte
	i := len(buf)
	for u > 0 {
		i--
		buf[i] = byte('0' + u%10)
		u /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
