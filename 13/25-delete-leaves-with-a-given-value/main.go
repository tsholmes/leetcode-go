package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var search func(*TreeNode) bool
	search = func(n *TreeNode) bool {
		if n == nil {
			return false
		}

		if search(n.Left) {
			n.Left = nil
		}
		if search(n.Right) {
			n.Right = nil
		}

		if n.Left == nil && n.Right == nil && n.Val == target {
			return true
		}

		return false
	}

	if search(root) {
		return nil
	}
	return root
}
