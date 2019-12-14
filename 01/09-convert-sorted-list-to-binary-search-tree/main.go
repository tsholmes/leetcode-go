package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	var build func(int, int) *TreeNode
	build = func(i, j int) *TreeNode {
		if i == j {
			return nil
		}
		mid := (i + j) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(i, mid),
			Right: build(mid+1, j),
		}
	}
	return build(0, len(nums))
}
