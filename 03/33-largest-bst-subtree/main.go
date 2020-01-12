package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestBSTSubtree(root *TreeNode) int {
	max := 0
	var search func(*TreeNode) (int, int, int)
	search = func(n *TreeNode) (int, int, int) {
		if n == nil {
			return 0, 0, 0
		}
		ok := true
		sz := 1
		minv, maxv := n.Val, n.Val
		if n.Left != nil {
			lsz, lmin, lmax := search(n.Left)
			if n.Left.Val >= n.Val || lmax >= n.Val || lsz == 0 {
				ok = false
			}
			sz += lsz
			minv = lmin
		}
		if n.Right != nil {
			rsz, rmin, rmax := search(n.Right)
			if n.Right.Val <= n.Val || rmin <= n.Val || rsz == 0 {
				ok = false
			}
			sz += rsz
			maxv = rmax
		}
		if ok {
			if sz > max {
				max = sz
			}
			return sz, minv, maxv
		}
		return 0, 0, 0
	}
	search(root)
	return max
}
