package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	seen := map[*TreeNode]bool{}
	add := func(n *TreeNode) {
		if !seen[n] {
			res = append(res, n.Val)
			seen[n] = true
		}
	}

	var walk func(*TreeNode, bool, bool)
	walk = func(n *TreeNode, isl bool, isr bool) {
		if n == nil {
			return
		}
		if isr {
			add(n)
		}
		walk(n.Right, isl && n.Left == nil, isr)
		walk(n.Left, isl, isr && n.Right == nil)
		if isl {
			add(n)
		}
		if n.Left == nil && n.Right == nil {
			add(n)
		}
	}
	walk(root.Right, false, true)
	walk(root.Left, true, false)
	add(root)

	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}

	return res
}
