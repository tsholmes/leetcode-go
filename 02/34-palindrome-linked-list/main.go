package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	valid := true
	var check func(*ListNode) *ListNode
	check = func(n *ListNode) *ListNode {
		if n == nil {
			return head
		}
		n2 := check(n.Next)
		if n.Val != n2.Val {
			valid = false
		}
		return n2.Next
	}
	check(head)
	return valid
}
