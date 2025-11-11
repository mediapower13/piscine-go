package piscine

// ListClear deletes all nodes from the linked list l.
// If l is nil, the function does nothing.
func ListClear(l *List) {
	if l == nil {
		return
	}
	// Reset the list to its zero value so Head and Tail become nil.
	*l = List{}
}
