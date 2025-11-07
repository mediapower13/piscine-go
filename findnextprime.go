package piscine

// FindNextPrime returns the first prime number that is equal or greater than nb.
// Only positive numbers can be prime; for nb <= 2 it returns 2.
func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	for {
		if IsPrime(nb) {
			return nb
		}
		nb += 2 // skip even numbers
	}
}
