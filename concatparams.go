package piscine

// ConcatParams concatenates the slice of arguments into a single string
// with a newline ('\n') between each argument. No trailing newline is added.
func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}
	res := args[0]
	for i := 1; i < len(args); i++ {
		res += "\n" + args[i]
	}
	return res
}
