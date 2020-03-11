package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var inorder func(*TreeNode, func(int))
	inorder = func(n *TreeNode, f func(int)) {
		if n == nil {
			return
		}
		inorder(n.Left, f)
		f(n.Val)
		inorder(n.Right, f)
	}

	counts := map[int]int{}
	inorder(root, func(i int) { counts[i]++ })

	max := 0
	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	var res []int
	for k, v := range counts {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}
