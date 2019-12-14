package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var search func(*TreeNode, int) bool
	search = func(node *TreeNode, s int) bool {
		s += node.Val
		if node.Left == nil && node.Right == nil {
			return s == sum
		}
		if node.Left != nil && search(node.Left, s) {
			return true
		}
		if node.Right != nil && search(node.Right, s) {
			return true
		}
		return false
	}
	return search(root, 0)
}
