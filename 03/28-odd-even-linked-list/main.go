package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var ohead, ehead *ListNode
	var otail, etail *ListNode
	for head != nil {
		n := head.Next
		if ohead == nil {
			ohead = head
			otail = head
		} else {
			otail.Next = head
			otail = head
		}
		head = n
		if head == nil {
			break
		}
		n = head.Next
		if ehead == nil {
			ehead = head
			etail = head
		} else {
			etail.Next = head
			etail = head
		}
		head = n
	}
	otail.Next = ehead
	if etail != nil {
		etail.Next = nil
	}
	return ohead
}
