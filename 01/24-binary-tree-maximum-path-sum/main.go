package main

import "fmt"

func main() {
	fmt.Println(maxPathSum(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))
	fmt.Println(maxPathSum(&TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	totalMax := root.Val
	var walk func(*TreeNode) int
	walk = func(node *TreeNode) int {
		max := node.Val
		lmax, rmax := 0, 0
		if node.Left != nil {
			lmax = walk(node.Left)
			if lmax+node.Val > max {
				max = lmax + node.Val
			}
		}
		if node.Right != nil {
			rmax = walk(node.Right)
		}
		if lmax+rmax+node.Val > totalMax {
			totalMax = lmax + rmax + node.Val
		}
		if lmax+node.Val > max {
			max = lmax + node.Val
		}
		if rmax+node.Val > max {
			max = rmax + node.Val
		}
		if max > totalMax {
			totalMax = max
		}
		if max > 0 {
			return max
		}
		return 0
	}
	walk(root)
	return totalMax
}
