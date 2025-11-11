package piscine

// BTreeTransplant replaces the subtree rooted at node with the subtree rooted
// at rplc in the tree. It returns the (possibly new) root of the tree.
// If node is nil the tree is unchanged. rplc may be nil (meaning removal).
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// If node is the root of the tree, rplc becomes the new root.
	if node.Parent == nil {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	parent := node.Parent
	if parent.Left == node {
		parent.Left = rplc
	} else {
		parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = parent
	}

	return root
}
