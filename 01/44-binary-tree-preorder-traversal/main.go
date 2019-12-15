package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var search func(*TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		search(node.Left)
		search(node.Right)
	}
	search(root)
	return res
}
