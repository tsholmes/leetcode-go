package main

import "fmt"

func main() {
	fmt.Println(flatten(reverseBetween(unflatten(1, 2, 3, 4, 5), 2, 4)))
	fmt.Println(flatten(reverseBetween(unflatten(1, 2, 3, 4, 5), 1, 5)))
	fmt.Println(flatten(reverseBetween(unflatten(1, 2, 3, 4, 5), 2, 6)))
	fmt.Println(flatten(reverseBetween(unflatten(1, 2, 3, 4, 5), 6, 7)))
	fmt.Println(flatten(reverseBetween(unflatten(), 1, 3)))
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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var preHead, preTail, revHead, revTail *ListNode

	for i := 1; i < m && head != nil; i++ {
		if preHead == nil {
			preHead = head
		}
		preTail = head
		head = head.Next
	}

	for i := m; i <= n && head != nil; i++ {
		node := head
		head = head.Next
		if revTail == nil {
			revTail = node
		}
		node.Next = revHead
		revHead = node
	}

	if preTail != nil {
		preTail.Next = revHead
	} else {
		preHead = revHead
	}
	if revTail != nil {
		revTail.Next = head
	}

	return preHead
}
