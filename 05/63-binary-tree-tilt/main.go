package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res := 0
	var search func(n *TreeNode) int
	search = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := search(n.Left), search(n.Right)
		res += abs(l - r)
		return l + r + n.Val
	}
	search(root)
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
