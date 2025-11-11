package piscine

// ListAt returns the pointer to the NodeL at position pos (0-based) starting from l.
// If l is nil, pos is negative, or pos is out of range, it returns nil.
func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	cur := l
	idx := 0
	for cur != nil {
		if idx == pos {
			return cur
		}
		cur = cur.Next
		idx++
	}
	return nil
}
