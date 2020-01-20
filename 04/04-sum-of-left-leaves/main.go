package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var search func(*TreeNode, bool)
	search = func(n *TreeNode, left bool) {
		if n == nil {
			return
		}
		search(n.Left, true)
		search(n.Right, false)
		if left && n.Left == nil && n.Right == nil {
			sum += n.Val
		}
	}
	search(root, false)
	return sum
}
