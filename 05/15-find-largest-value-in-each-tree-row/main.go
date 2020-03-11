package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	var search func(*TreeNode, int)
	search = func(n *TreeNode, d int) {
		if n == nil {
			return
		}
		if d == len(res) {
			res = append(res, n.Val)
		} else if res[d] < n.Val {
			res[d] = n.Val
		}
		search(n.Left, d+1)
		search(n.Right, d+1)
	}
	search(root, 0)
	return res
}
