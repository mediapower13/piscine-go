package piscine

// BTreeApplyPreorder applies function f to each node's Data following preorder
// traversal (Node, Left, Right). f is a variadic function matching
// fmt.Println-like signature (func(...interface{}) (int, error)).
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	// call f with the node data; ignore returned values
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
