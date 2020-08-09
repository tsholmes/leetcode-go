package main

import "strconv"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	var res []byte

	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = strconv.AppendInt(res, int64(n.Val), 10)
		if n.Left == nil && n.Right == nil {
			return
		}
		res = append(res, '(')
		walk(n.Left)
		res = append(res, ')')
		if n.Right != nil {
			res = append(res, '(')
			walk(n.Right)
			res = append(res, ')')
		}
	}

	walk(t)

	return string(res)
}
