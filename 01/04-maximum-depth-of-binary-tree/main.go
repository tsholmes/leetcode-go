package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d := maxDepth(root.Left)
	if d2 := maxDepth(root.Right); d2 > d {
		d = d2
	}
	return d + 1
}
