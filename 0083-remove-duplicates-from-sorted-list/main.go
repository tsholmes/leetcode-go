package main

import "fmt"

func main() {
	fmt.Println(flatten(deleteDuplicates(unflatten(1, 1, 2))))
	fmt.Println(flatten(deleteDuplicates(unflatten(1, 1, 2, 3, 3))))
	fmt.Println(flatten(deleteDuplicates(unflatten())))
	fmt.Println(flatten(deleteDuplicates(unflatten(1, 1))))
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead, newTail := head, head
	head = head.Next
	newTail.Next = nil
	for head != nil {
		n := head
		head = head.Next
		if n.Val != newTail.Val {
			newTail.Next = n
			newTail = newTail.Next
			newTail.Next = nil
		}
	}
	return newHead
}
