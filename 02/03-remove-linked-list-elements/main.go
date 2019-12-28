package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var rhead, rtail *ListNode
	for head != nil {
		if head.Val != val {
			if rhead == nil {
				rhead, rtail = head, head
			} else {
				rtail.Next = head
				rtail = head
			}
		}
		head = head.Next
	}
	if rtail != nil {
		rtail.Next = nil
	}
	return rhead
}
