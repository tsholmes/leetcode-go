package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	count := 0
	sums := []int{0}
	var search func(n *TreeNode, csum int)
	search = func(n *TreeNode, csum int) {
		if n == nil {
			return
		}
		slen := len(sums)
		csum += n.Val
		for _, s := range sums {
			if csum-s == sum {
				count++
			}
		}
		sums = append(sums, csum)
		search(n.Left, csum)
		search(n.Right, csum)
		sums = sums[:slen]
	}

	search(root, 0)
	return count
}
