package main

import "fmt"

func main() {
	fmt.Println(flatten(insertionSortList(unflatten(4, 2, 1, 3))))
	fmt.Println(flatten(insertionSortList(unflatten(-1, 5, 3, 4, 0))))
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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	shead := head
	uhead := head.Next
	shead.Next = nil
	for uhead != nil {
		n := uhead
		uhead = uhead.Next

		a, b := shead, shead.Next
		for b != nil && n.Val > b.Val {
			a, b = a.Next, b.Next
		}
		if a == shead && a.Val > n.Val {
			n.Next = shead
			shead = n
		} else {
			n.Next = a.Next
			a.Next = n
		}
	}
	return shead
}
