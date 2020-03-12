package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var search func(*TreeNode) int
	search = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := search(n.Left), search(n.Right)
		if l+r > res {
			res = l + r
		}
		if l > r {
			return 1 + l
		}
		return 1 + r
	}
	search(root)
	return res
}
