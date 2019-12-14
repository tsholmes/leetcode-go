package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	tail := root
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		l, r := node.Left, node.Right
		node.Left, node.Right = nil, nil
		tail.Right = node
		tail = node
		walk(l)
		walk(r)
	}
	l, r := root.Left, root.Right
	root.Left, root.Right = nil, nil
	walk(l)
	walk(r)
}
