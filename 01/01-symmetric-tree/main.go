package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var walk func(a, b *TreeNode) bool
	walk = func(a, b *TreeNode) bool {
		if a == nil {
			return b == nil
		}
		if b == nil {
			return false
		}
		return a.Val == b.Val && walk(a.Left, b.Right) && walk(a.Right, b.Left)
	}
	return walk(root.Left, root.Right)
}
