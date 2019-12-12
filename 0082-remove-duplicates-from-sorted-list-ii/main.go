package main

import "fmt"

func main() {
	fmt.Println(flatten(deleteDuplicates(unflatten(1, 2, 3, 3, 4, 4, 5))))
	fmt.Println(flatten(deleteDuplicates(unflatten(1, 1, 1, 2, 3))))
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
	var newHead, newTail *ListNode
	for head != nil {
		if head.Next != nil && head.Next.Val == head.Val {
			v := head.Val
			for head != nil && head.Val == v {
				head = head.Next
			}
		} else {
			var n *ListNode
			n, head = head, head.Next
			n.Next = nil
			if newTail == nil {
				newHead, newTail = n, n
			} else {
				newTail.Next = n
				newTail = n
			}
		}
	}
	return newHead
}
