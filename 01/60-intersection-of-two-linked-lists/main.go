package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA, hB := headA, headB
	ac, bc := 0, 0
	for hA != nil {
		ac++
		hA = hA.Next
	}
	for hB != nil {
		bc++
		hB = hB.Next
	}
	for i := 0; i < ac-bc; i++ {
		headA = headA.Next
	}
	for j := 0; j < bc-ac; j++ {
		headB = headB.Next
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
