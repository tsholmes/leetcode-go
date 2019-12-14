package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	work := []*TreeNode{root}
	var next []*TreeNode

	forward := true

	for len(work) > 0 {
		var row []int
		for _, n := range work {
			row = append(row, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if !forward {
			for i := 0; i < len(row)/2; i++ {
				row[i], row[len(row)-i-1] = row[len(row)-i-1], row[i]
			}
		}
		res = append(res, row)
		work, next = next, work[:0]
		forward = !forward
	}

	return res
}
