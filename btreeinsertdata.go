package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node containing data into the binary search tree
// that has root `root`. It returns the (possibly new) root of the tree.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	cur := root
	var parent *TreeNode

	for cur != nil {
		parent = cur
		if data < cur.Data {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	newNode := &TreeNode{Data: data, Parent: parent}
	if data < parent.Data {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	return root
}
