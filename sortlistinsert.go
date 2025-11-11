package piscine

// SortListInsert inserts a new NodeI with Data=data_ref into the sorted linked list l
// preserving ascending order. It returns the new head of the list.
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil {
		return newNode
	}
	// Insert at head if needed
	if data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	cur := l
	for cur.Next != nil && cur.Next.Data < data_ref {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	return l
}
