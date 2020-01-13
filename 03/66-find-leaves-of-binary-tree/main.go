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

func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	at := map[int][]int{}
	var search func(*TreeNode) int
	search = func(n *TreeNode) int {
		d := 0
		if n.Left != nil {
			d = search(n.Left) + 1
		}
		if n.Right != nil {
			d2 := search(n.Right) + 1
			if d2 > d {
				d = d2
			}
		}
		at[d] = append(at[d], n.Val)
		return d
	}
	search(root)
	var ks []int
	for k := range at {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	var res [][]int
	for _, k := range ks {
		res = append(res, at[k])
	}
	return res
}
