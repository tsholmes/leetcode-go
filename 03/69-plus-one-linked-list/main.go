package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	var search func(*ListNode) bool
	search = func(n *ListNode) bool {
		carry := false
		if n.Next == nil || search(n.Next) {
			n.Val++
		}
		if n.Val >= 10 {
			n.Val -= 10
			carry = true
		}
		return carry
	}
	if search(head) {
		head = &ListNode{
			Val:  1,
			Next: head,
		}
	}
	return head
}
