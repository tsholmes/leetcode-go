package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var work []int

	var search func(*TreeNode, int)
	search = func(node *TreeNode, s int) {
		wl := len(work)
		defer func() {
			work = work[:wl]
		}()

		s += node.Val
		work = append(work, node.Val)
		if node.Left == nil && node.Right == nil {
			if s == sum {
				res = append(res, append([]int{}, work...))
			}
			return
		}
		if node.Left != nil {
			search(node.Left, s)
		}
		if node.Right != nil {
			search(node.Right, s)
		}
	}
	search(root, 0)
	return res
}
