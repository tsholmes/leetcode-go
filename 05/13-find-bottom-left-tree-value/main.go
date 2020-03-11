package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	maxd := -1
	res := 0
	var search func(*TreeNode, int)
	search = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		if depth > maxd {
			maxd = depth
			res = n.Val
		}
		search(n.Left, depth+1)
		search(n.Right, depth+1)
	}
	search(root, 0)
	return res
}
