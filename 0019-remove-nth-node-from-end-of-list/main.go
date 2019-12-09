package main

import "fmt"

func main() {
	fmt.Println(flatten(removeNthFromEnd(unflatten(1, 2, 3, 4, 5), 2)))
	fmt.Println(flatten(removeNthFromEnd(unflatten(1, 2, 3, 4, 5), 1)))
	fmt.Println(flatten(removeNthFromEnd(unflatten(1, 2, 3, 4, 5), 5)))
	fmt.Println(flatten(removeNthFromEnd(unflatten(1), 1)))
}

func flatten(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func unflatten(list ...int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	return &ListNode{
		Val:  list[0],
		Next: unflatten(list[1:]...),
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fwd := head
	for i := 0; i < n-1; i++ {
		fwd = fwd.Next
	}
	if fwd.Next == nil {
		return head.Next
	}
	back := head
	for {
		fwd = fwd.Next
		if fwd.Next == nil {
			back.Next = back.Next.Next
			break
		}
		back = back.Next
	}
	return head
}
