package main

import "sort"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var res []int
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		walk(n.Left)
		walk(n.Right)
	}
	walk(root1)
	walk(root2)
	sort.Ints(res)
	return res
}
