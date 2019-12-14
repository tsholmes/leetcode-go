package main

import "fmt"

func main() {
	fmt.Println(isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
	fmt.Println(isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var search func(*TreeNode) (int, int, bool)
	search = func(node *TreeNode) (int, int, bool) {
		min, max := node.Val, node.Val
		var ok bool
		if node.Left != nil {
			var lmax int
			min, lmax, ok = search(node.Left)
			if !ok {
				return 0, 0, false
			}
			if lmax >= node.Val {
				return 0, 0, false
			}
		}
		if node.Right != nil {
			var rmin int
			rmin, max, ok = search(node.Right)
			if !ok {
				return 0, 0, false
			}
			if rmin <= node.Val {
				return 0, 0, false
			}
		}
		return min, max, true
	}
	_, _, ok := search(root)
	return ok
}
