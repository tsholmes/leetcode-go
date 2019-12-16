package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var right, left *TreeNode
	for root.Left != nil {
		left, right, root.Left, root.Right, root = root.Right, root, left, right, root.Left
	}
	root.Left = left
	root.Right = right
	return root
}
