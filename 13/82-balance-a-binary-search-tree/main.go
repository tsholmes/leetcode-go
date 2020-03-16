package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var nums []int
	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		nums = append(nums, n.Val)
		inorder(n.Right)
	}
	inorder(root)

	var build func(i, j int) *TreeNode
	build = func(i, j int) *TreeNode {
		if i > j {
			return nil
		}
		if i == j {
			return &TreeNode{
				Val: nums[i],
			}
		}
		mid := (i + j) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(i, mid-1),
			Right: build(mid+1, j),
		}
	}

	return build(0, len(nums)-1)
}
