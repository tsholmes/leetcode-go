package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var exact func(a, b *TreeNode) bool
	exact = func(a, b *TreeNode) bool {
		if (a == nil) != (b == nil) {
			return false
		}
		if a == nil {
			return true
		}
		if a.Val != b.Val {
			return false
		}
		return exact(a.Left, b.Left) && exact(a.Right, b.Right)
	}

	var search func(a *TreeNode) bool
	search = func(a *TreeNode) bool {
		if a == nil {
			return t == nil
		}
		return exact(a, t) || search(a.Left) || search(a.Right)
	}

	return search(s)
}
