package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	var lte []int
	var gt []int
	var search func(*TreeNode)
	search = func(n *TreeNode) {
		if n == nil {
			return
		}
		f := float64(n.Val)
		if f <= target {
			search(n.Right)
			if len(lte) < k {
				lte = append(lte, n.Val)
			}
			if len(lte) < k {
				search(n.Left)
			}
		} else {
			search(n.Left)
			if len(gt) < k {
				gt = append(gt, n.Val)
			}
			if len(gt) < k {
				search(n.Right)
			}
		}
	}
	search(root)

	for i := 0; i < len(lte)/2; i++ {
		j := len(lte) - i - 1
		lte[i], lte[j] = lte[j], lte[i]
	}

	res := append(lte, gt...)
	for len(res) > k {
		if target-float64(res[0]) > float64(res[len(res)-1])-target {
			res = res[1:]
		} else {
			res = res[:len(res)-1]
		}
	}
	return res
}
