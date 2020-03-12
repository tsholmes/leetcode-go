package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var revorder func(*TreeNode, func(*TreeNode))
	revorder = func(n *TreeNode, f func(*TreeNode)) {
		if n == nil {
			return
		}
		revorder(n.Right, f)
		f(n)
		revorder(n.Left, f)
	}
	sum := 0
	revorder(root, func(n *TreeNode) {
		sum += n.Val
		n.Val = sum
	})
	return root
}
