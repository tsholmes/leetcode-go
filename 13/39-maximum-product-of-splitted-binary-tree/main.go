package main

import "fmt"

func main() {
	fmt.Println(maxProduct(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	tsum := 0
	var search func(n *TreeNode)
	search = func(n *TreeNode) {
		if n == nil {
			return
		}
		tsum += n.Val
		search(n.Left)
		search(n.Right)
	}
	search(root)

	max := 0
	var ssearch func(n *TreeNode) int
	ssearch = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		lsum := ssearch(n.Left)
		rsum := ssearch(n.Right)
		if n.Left != nil && lsum*(tsum-lsum) > max {
			max = lsum * (tsum - lsum)
		}
		if n.Right != nil && rsum*(tsum-rsum) > max {
			max = rsum * (tsum - rsum)
		}
		return n.Val + lsum + rsum
	}
	ssearch(root)
	return max % (1e9 + 7)
}
