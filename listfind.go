package piscine

// CompStr compares two interface{} values using ==.
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind returns the address of the Data field of the first node
// where comp(node.Data, ref) == true. Returns nil if not found.
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l == nil || comp == nil {
		return nil
	}
	for node := l.Head; node != nil; node = node.Next {
		if comp(node.Data, ref) {
			return &node.Data
		}
	}
	return nil
}
