package piscine

// SortedListMerge merges two sorted linked lists n1 and n2 into a single
// sorted list and returns its head. Nodes are reused; no new nodes are created.
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	var head *NodeI
	// initialize head
	if n1.Data <= n2.Data {
		head = n1
		n1 = n1.Next
	} else {
		head = n2
		n2 = n2.Next
	}

	cur := head
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			cur.Next = n1
			n1 = n1.Next
		} else {
			cur.Next = n2
			n2 = n2.Next
		}
		cur = cur.Next
	}

	if n1 != nil {
		cur.Next = n1
	} else {
		cur.Next = n2
	}

	return head
}
