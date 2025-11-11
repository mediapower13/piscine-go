package piscine

// BTreeDeleteNode deletes node from the BST rooted at root and returns the
// possibly-updated root. It uses BTreeTransplant/BTreeMin helpers.
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case: node has no left child -> replace node with its right child
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
		return root
	}

	// Case: node has no right child -> replace node with its left child
	if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
		return root
	}

	// Node has two children: find its successor (minimum in right subtree)
	succ := BTreeMin(node.Right)
	if succ == nil {
		// should not happen because node.Right != nil
		return root
	}

	if succ.Parent != node {
		// Replace successor with its right child
		root = BTreeTransplant(root, succ, succ.Right)
		// Attach node's right subtree to successor
		succ.Right = node.Right
		if succ.Right != nil {
			succ.Right.Parent = succ
		}
	}

	// Replace node with successor
	root = BTreeTransplant(root, node, succ)
	// Attach node's left subtree to successor
	succ.Left = node.Left
	if succ.Left != nil {
		succ.Left.Parent = succ
	}

	return root
}
