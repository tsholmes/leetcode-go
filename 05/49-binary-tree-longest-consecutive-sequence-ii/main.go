package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	max := 0
	var search func(n *TreeNode) (int, int)
	search = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		lu, ld := search(n.Left)
		ru, rd := search(n.Right)

		maxu, maxd := 1, 1
		if n.Left != nil {
			if n.Left.Val == n.Val-1 {
				maxu = 1 + lu
			}
			if n.Left.Val == n.Val+1 {
				maxd = 1 + ld
			}
		}
		if n.Right != nil {
			if n.Right.Val == n.Val-1 && ru >= maxu {
				maxu = 1 + ru
			}
			if n.Right.Val == n.Val+1 && rd >= maxd {
				maxd = 1 + rd
			}
		}

		if maxu+maxd-1 > max {
			max = maxu + maxd - 1
		}

		return maxu, maxd
	}
	search(root)
	return max
}
