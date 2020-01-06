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

func longestConsecutive(root *TreeNode) int {
	max := 0
	var search func(*TreeNode, int, int)
	search = func(n *TreeNode, last int, l int) {
		if n == nil {
			return
		}
		if n.Val == last+1 {
			l++
		} else {
			l = 1
		}
		if l > max {
			max = l
		}
		search(n.Left, n.Val, l)
		search(n.Right, n.Val, l)
	}
	search(root, math.MinInt64, 0)
	return max
}
