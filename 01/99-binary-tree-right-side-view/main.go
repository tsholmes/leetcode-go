package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	var walk func(*TreeNode, int)
	walk = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		if depth == len(res) {
			res = append(res, n.Val)
		} else {
			res[depth] = n.Val
		}

		walk(n.Left, depth+1)
		walk(n.Right, depth+1)
	}
	walk(root, 0)
	return res
}
