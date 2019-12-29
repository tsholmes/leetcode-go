package main

func main() {

}

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var search func(*TreeNode) *TreeNode
	search = func(n *TreeNode) *TreeNode {
		if n == nil || n == p || n == q {
			return n
		}
		a, b := search(n.Left), search(n.Right)
		if a != nil && b != nil {
			return n
		} else if a != nil {
			return a
		} else {
			return b
		}
	}
	return search(root)
}
