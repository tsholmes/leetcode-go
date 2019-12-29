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
	if p.Val > q.Val {
		p, q = q, p
	}
	var search func(*TreeNode) *TreeNode
	search = func(n *TreeNode) *TreeNode {
		if p.Val <= n.Val && q.Val >= n.Val {
			return n
		}
		if p.Val > n.Val {
			return search(n.Right)
		} else {
			return search(n.Left)
		}
	}
	return search(root)
}
