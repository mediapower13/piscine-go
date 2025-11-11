package piscine

// BTreeMin returns the node with the minimum Data value in the BST rooted at
// root. If root is nil, it returns nil.
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}
