package piscine

// IsPrime returns true if nb is a prime number, false otherwise.
// Only positive integers > 1 can be prime.
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
