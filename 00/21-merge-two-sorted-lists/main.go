package main

import "fmt"

func main() {
	fmt.Println(flatten(mergeTwoLists(unflatten(1, 2, 4), unflatten(1, 3, 4))))
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := l1
	if l2.Val < l1.Val {
		head = l2
		l2 = l2.Next
	} else {
		l1 = l1.Next
	}
	tail := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}

	return head
}
