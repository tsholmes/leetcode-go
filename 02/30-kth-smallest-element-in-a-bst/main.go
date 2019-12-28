package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var (
		found bool
		res   int
		count int
		walk  func(*TreeNode)
	)
	walk = func(n *TreeNode) {
		if found || n == nil {
			return
		}
		walk(n.Left)
		count++
		if count == k {
			res = n.Val
			found = true
			return
		}
		walk(n.Right)
	}
	walk(root)
	return res
}
