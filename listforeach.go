package piscine

// ListForEach applies function f to each node in list l.
// If l or f is nil the function does nothing.
func ListForEach(l *List, f func(*NodeL)) {
	if l == nil || f == nil {
		return
	}
	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}

// Add2_node adds 2 to integer nodes or appends "2" to string nodes.
func Add2_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v + 2
	case string:
		node.Data = v + "2"
	}
}

// Subtract3_node subtracts 3 from integer nodes or appends "-3" to string nodes.
func Subtract3_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v - 3
	case string:
		node.Data = v + "-3"
	}
}
