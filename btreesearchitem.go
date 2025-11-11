package piscine

// BTreeSearchItem searches for a node with Data equal to elem in the BST rooted
// at root and returns the node pointer if found, otherwise nil.
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
    cur := root
    for cur != nil {
        if elem == cur.Data {
            return cur
        }
        if elem < cur.Data {
            cur = cur.Left
        } else {
            cur = cur.Right
        }
    }
    return nil
}
