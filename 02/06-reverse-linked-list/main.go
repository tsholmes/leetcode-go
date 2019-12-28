package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rhead, rtail *ListNode
	var walk func(*ListNode)
	walk = func(n *ListNode) {
		if n == nil {
			return
		}
		walk(n.Next)
		if rhead == nil {
			rhead = n
			rtail = n
		} else {
			rtail.Next = n
			rtail = n
		}
	}
	walk(head)
	if rtail != nil {
		rtail.Next = nil
	}
	return rhead
}
