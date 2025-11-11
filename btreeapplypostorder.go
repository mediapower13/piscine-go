package piscine

// BTreeApplyPostorder applies function f to each node's Data following postorder
// traversal (Left, Right, Node). f is a variadic function matching
// fmt.Println-like signature (func(...interface{}) (int, error)).
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	// call f with the node data; ignore returned values
	f(root.Data)
}
