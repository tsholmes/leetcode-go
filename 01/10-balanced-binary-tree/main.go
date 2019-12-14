package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var search func(*TreeNode) (int, bool)
	search = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		dl, lok := search(node.Left)
		dr, rok := search(node.Right)
		if !lok || !rok {
			return 0, false
		}
		dd := dl - dr
		if dd < -1 || dd > 1 {
			return 0, false
		}
		d := dl
		if dr > d {
			d = dr
		}
		return d + 1, true
	}
	_, ok := search(root)
	return ok
}
