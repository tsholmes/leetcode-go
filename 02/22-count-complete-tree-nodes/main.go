package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0
	var search func(*TreeNode)
	search = func(n *TreeNode) {
		if n == nil {
			return
		}
		count++
		search(n.Left)
		search(n.Right)
	}
	search(root)
	return count
}
