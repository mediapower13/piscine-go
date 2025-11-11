package piscine

// NodeI is a singly-linked list node holding an int value.
type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the linked list of NodeI in ascending order and returns the head.
// It sorts in-place by swapping Data values (no new nodes created).
func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	// Bubble sort on linked list
	swapped := true
	for swapped {
		swapped = false
		cur := l
		for cur != nil && cur.Next != nil {
			if cur.Data > cur.Next.Data {
				cur.Data, cur.Next.Data = cur.Next.Data, cur.Data
				swapped = true
			}
			cur = cur.Next
		}
	}
	return l
}
