package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var walk func(*TreeNode, *TreeNode) bool
	walk = func(a, b *TreeNode) bool {
		if a == nil {
			return b == nil
		}
		if b == nil {
			return false
		}
		return a.Val == b.Val && walk(a.Left, b.Left) && walk(a.Right, b.Right)
	}
	return walk(p, q)
}
