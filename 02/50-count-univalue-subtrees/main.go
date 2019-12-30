package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var walk func(*TreeNode) (int, bool)
	walk = func(n *TreeNode) (int, bool) {
		res := 0
		valid := true
		if n.Left != nil {
			c, ok := walk(n.Left)
			res += c
			valid = valid && ok && n.Val == n.Left.Val
		}
		if n.Right != nil {
			c, ok := walk(n.Right)
			res += c
			valid = valid && ok && n.Val == n.Right.Val
		}
		if valid {
			res++
		}
		return res, valid
	}
	res, _ := walk(root)
	return res
}
