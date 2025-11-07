package piscine

// Sqrt returns the integer square root of nb when it is a perfect square.
// If nb is negative or does not have an integer square root, returns 0.
func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}

	// binary search between 1 and nb/2
	lo := 1
	hi := nb / 2
	for lo <= hi {
		mid := (lo + hi) / 2
		sq := mid * mid
		if sq == nb {
			return mid
		}
		if sq < nb {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return 0
}
