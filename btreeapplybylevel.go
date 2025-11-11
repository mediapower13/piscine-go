package piscine

// BTreeApplyByLevel applies function f to each node's Data following a
// level-order (breadth-first) traversal. f is a variadic function matching
// fmt.Println-like signature (func(...interface{}) (int, error)).
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		f(node.Data)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
