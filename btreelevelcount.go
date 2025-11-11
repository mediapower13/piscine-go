package piscine

// BTreeLevelCount returns the number of levels in the binary tree (height).
// An empty tree has 0 levels; a single node tree has 1.
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := BTreeLevelCount(root.Left)
	right := BTreeLevelCount(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
