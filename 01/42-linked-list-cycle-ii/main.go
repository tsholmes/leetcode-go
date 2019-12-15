package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	h1, h2 := head, head
	for h1 != nil && h2 != nil {
		h1 = h1.Next
		h2 = h2.Next
		if h2 == nil {
			return nil
		}
		h2 = h2.Next
		if h1 == h2 {
			break
		}
	}
	if h1 == nil || h2 == nil {
		return nil
	}
	for {
		h1 = h1.Next
		h2 = h2.Next.Next
		if h1 == head || h2 == head {
			return head
		}
		if h1 == h2 {
			head = head.Next
		}
	}
}
