package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var nodes []*TreeNode
	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		if n.Val != key {
			nodes = append(nodes, n)
		}
		walk(n.Right)
	}
	walk(root)
	if len(nodes) == 0 {
		return nil
	}
	root = nodes[0]
	right := root
	root.Left, root.Right = nil, nil
	for _, n := range nodes[1:] {
		n.Left, n.Right = nil, nil
		right.Right, right = n, n
	}
	return root
}
