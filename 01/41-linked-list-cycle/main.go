package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	h1, h2 := head, head
	for h1 != nil && h2 != nil {
		h1 = h1.Next
		h2 = h2.Next
		if h2 == nil {
			return false
		}
		h2 = h2.Next
		if h1 == h2 {
			return true
		}
	}
	return false
}
