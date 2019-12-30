package main

import "strconv"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	var work []byte
	var search func(*TreeNode)
	search = func(n *TreeNode) {
		if len(work) > 0 {
			work = append(work, '-', '>')
		}
		work = strconv.AppendInt(work, int64(n.Val), 10)
		if n.Left == nil && n.Right == nil {
			res = append(res, string(work))
			return
		}
		wlen := len(work)
		if n.Left != nil {
			search(n.Left)
			work = work[:wlen]
		}
		if n.Right != nil {
			search(n.Right)
			work = work[:wlen]
		}
	}
	search(root)
	return res
}
