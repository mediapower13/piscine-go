package piscine

// Fibonacci returns the value at position index in the Fibonacci sequence.
// The first value is at index 0. A negative index returns -1.
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
