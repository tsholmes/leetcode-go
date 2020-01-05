package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	var search func(*TreeNode)
	search = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Val > p.Val {
			res = n
			search(n.Left)
		} else {
			search(n.Right)
		}
	}
	search(root)
	return res
}
