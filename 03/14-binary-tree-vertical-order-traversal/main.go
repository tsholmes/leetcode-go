package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(verticalOrder(&TreeNode{
		Val:  3,
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

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	type pair struct {
		num   int
		level int
	}
	minOff := 0
	maxOff := 0
	offsets := map[int][]pair{}

	var search func(*TreeNode, int, int)
	search = func(n *TreeNode, level int, offset int) {
		if n == nil {
			return
		}
		offsets[offset] = append(offsets[offset], pair{n.Val, level})
		if offset < minOff {
			minOff = offset
		}
		if offset > maxOff {
			maxOff = offset
		}
		search(n.Left, level+1, offset-1)
		search(n.Right, level+1, offset+1)
	}

	search(root, 0, 0)

	var res [][]int

	for off := minOff; off <= maxOff; off++ {
		ps := offsets[off]
		sort.SliceStable(ps, func(i, j int) bool {
			pi, pj := ps[i], ps[j]
			return pi.level < pj.level
		})
		r := make([]int, len(ps))
		for i, p := range ps {
			r[i] = p.num
		}
		res = append(res, r)
	}

	return res
}
