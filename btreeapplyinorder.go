package piscine

// BTreeApplyInorder applies function f to each node's Data following inorder
// traversal (Left, Node, Right). f is a variadic function matching
// fmt.Println-like signature (func(...interface{}) (int, error)).
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	// call f with the node data; ignore returned values
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
