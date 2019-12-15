package main

import "fmt"

func main() {
	fmt.Println(flatten(sortList(unflatten(4, 2, 1, 3))))
	fmt.Println(flatten(sortList(unflatten(-1, 5, 3, 4, 0))))
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

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	n0, n1, n2 := head, head, head
	for n1 != nil && n2 != nil {
		if n1 != head {
			n0 = n1
		}
		n1 = n1.Next
		n2 = n2.Next
		if n2 != nil {
			n2 = n2.Next
		}
	}

	n0.Next = nil

	return mergeTwoLists(sortList(head), sortList(n1))
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
