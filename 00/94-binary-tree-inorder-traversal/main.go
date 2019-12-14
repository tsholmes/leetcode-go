package main

import "fmt"

func main() {
	fmt.Println(inorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
	fmt.Println(inorderTraversal(nil))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		res = append(res, node.Val)
		walk(node.Right)
	}
	walk(root)
	return res
}
