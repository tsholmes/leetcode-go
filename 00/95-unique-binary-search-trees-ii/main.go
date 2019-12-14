package main

import (
	"fmt"
	"strconv"
)

func main() {
	do := func(n int) {
		res := generateTrees(n)
		for _, t := range res {
			fmt.Println(flatten(t))
		}
		fmt.Println()
	}
	do(1)
	do(3)
	do(5)
}

func flatten(root *TreeNode) []string {
	var res []string
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			res = append(res, "nil")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			if node.Left != nil || node.Right != nil {
				walk(node.Left)
			}
			if node.Right != nil {
				walk(node.Right)
			}
		}
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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	memo := map[[2]int][]*TreeNode{}
	empty := []*TreeNode{nil}
	var search func(int, int) []*TreeNode
	search = func(i, j int) []*TreeNode {
		if i > j {
			return empty
		}
		key := [2]int{i, j}
		res, ok := memo[key]
		if ok {
			return res
		}
		for c := i; c <= j; c++ {
			ls, rs := search(i, c-1), search(c+1, j)
			for _, l := range ls {
				for _, r := range rs {
					res = append(res, &TreeNode{
						Val:   c,
						Left:  l,
						Right: r,
					})
				}
			}
		}
		memo[key] = res
		return res
	}
	return search(1, n)
}
