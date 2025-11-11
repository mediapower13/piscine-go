package piscine

// ListMerge appends all nodes from l2 to the end of l1 without creating new nodes.
// After the operation l2 will be empty (Head and Tail set to nil).
func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l2 == nil {
		return
	}

	if l2.Head == nil {
		// nothing to merge
		return
	}

	if l1.Head == nil {
		// l1 is empty: take l2 entirely
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		l2.Head = nil
		l2.Tail = nil
		return
	}

	// both non-empty: link l1.Tail to l2.Head
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail

	// empty l2
	l2.Head = nil
	l2.Tail = nil
}
