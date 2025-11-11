package piscine

// BTreeMax returns the node with the maximum Data value in the BST rooted at
// root. If root is nil, it returns nil.
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}
