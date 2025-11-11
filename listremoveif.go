package piscine

// ListRemoveIf removes all nodes whose Data equals data_ref from the list l.
// It updates Head and Tail accordingly. If l is nil, it does nothing.
func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	// Remove matching nodes at the head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If list became empty, reset Tail and return
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Now remove matching nodes in the rest of the list
	prev := l.Head
	cur := l.Head.Next
	for cur != nil {
		if cur.Data == data_ref {
			// remove cur
			prev.Next = cur.Next
			if cur == l.Tail {
				l.Tail = prev
			}
			cur = prev.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}
}
