package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var search func(*TreeNode) (int, int)
	search = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		l1, l2 := search(n.Left)
		r1, r2 := search(n.Right)

		return n.Val + l2 + r2, max(l1, l2) + max(r1, r2)
	}
	return max(search(root))
}
