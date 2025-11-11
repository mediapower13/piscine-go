package piscine

// ListReverse reverses the order of elements in the linked list l.
// It updates both Head and Tail. If l is nil or empty, it does nothing.
func ListReverse(l *List) {
    if l == nil || l.Head == nil {
        return
    }

    var prev *NodeL
    cur := l.Head
    // After reversal the original head becomes the tail
    l.Tail = l.Head

    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    l.Head = prev
}
