package piscine

func ListSize(l *List) int {
	if l == nil || l.Head == nil {
		return 0
	}
	count := 0
	for node := l.Head; node != nil; node = node.Next {
		count++
	}
	return count
}
