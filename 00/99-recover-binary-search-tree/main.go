package main

import (
	"fmt"
	"sort"
)

func main() {
	do := func(root *TreeNode) {
		recoverTree(root)
		fmt.Println(flatten(root))
	}

	do(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	})
	do(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
	})
	do(nil)
}

func flatten(root *TreeNode) []int {
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

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var nodes []*TreeNode
	var vals []int
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		nodes = append(nodes, node)
		vals = append(vals, node.Val)
		walk(node.Right)
	}
	walk(root)

	sort.Ints(vals)
	for i := range nodes {
		nodes[i].Val = vals[i]
	}
}
