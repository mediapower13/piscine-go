package piscine

// ListLast returns the Data of the last element of the list l.
// If l is nil or empty, it returns nil.
func ListLast(l *List) interface{} {
    if l == nil || l.Head == nil {
        return nil
    }
    cur := l.Head
    for cur.Next != nil {
        cur = cur.Next
    }
    return cur.Data
}
