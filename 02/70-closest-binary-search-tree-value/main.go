package main

import "math"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest := 0
	diff := math.MaxFloat64

	for root != nil {
		d := float64(root.Val) - target
		if d < 0 {
			d = -d
		}
		if d < diff {
			diff = d
			closest = root.Val
		}
		if float64(root.Val) >= target {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}
