package main

import "math"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var inorder func(*TreeNode, func(int))
	inorder = func(n *TreeNode, f func(int)) {
		if n == nil {
			return
		}
		inorder(n.Left, f)
		f(n.Val)
		inorder(n.Right, f)
	}
	min := math.MaxInt32
	last := math.MinInt32
	inorder(root, func(v int) {
		if v-last < min {
			min = v - last
		}
		last = v
	})
	return min
}
