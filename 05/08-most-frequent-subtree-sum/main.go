package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	counts := map[int]int{}
	var search func(n *TreeNode) int
	search = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		sum := n.Val + search(n.Left) + search(n.Right)
		counts[sum]++
		return sum
	}
	search(root)
	max := -1
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
