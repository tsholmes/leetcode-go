package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	work := []*TreeNode{root}
	var next []*TreeNode

	for len(work) > 0 {
		var row []int
		for _, n := range work {
			if n != nil {
				row = append(row, n.Val)
				if n.Left != nil {
					next = append(next, n.Left)
				}
				if n.Right != nil {
					next = append(next, n.Right)
				}
			}
		}
		res = append(res, row)
		work, next = next, work[:0]
	}

	return res
}
