package piscine

// BTreeIsBinary returns true if the tree rooted at root satisfies the
// binary-search-tree property. It allows equal values on the right subtree
// (i.e. inorder traversal must be non-decreasing).
func BTreeIsBinary(root *TreeNode) bool {
	var prevVal string
	var prevSet bool
	ok := true

	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil || !ok {
			return
		}
		inorder(n.Left)
		if !ok {
			return
		}
		if prevSet {
			if prevVal > n.Data {
				ok = false
				return
			}
		}
		prevVal = n.Data
		prevSet = true
		inorder(n.Right)
	}

	inorder(root)
	return ok
}
