package main

import "fmt"

func main() {
	fmt.Println(sumNumbers(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))
	fmt.Println(sumNumbers(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   9,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var search func(*TreeNode, int)
	search = func(node *TreeNode, v int) {
		v = v*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += v
			return
		}
		if node.Left != nil {
			search(node.Left, v)
		}
		if node.Right != nil {
			search(node.Right, v)
		}
	}
	search(root, 0)
	return res
}
