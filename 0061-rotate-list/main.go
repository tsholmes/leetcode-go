package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(flatten(rotateRight(unflatten(1, 2, 3, 4, 5), i)))
	}
	fmt.Println(flatten(rotateRight(unflatten(0, 1, 2), 4)))
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	count := 1
	tail := head
	for tail.Next != nil {
		count++
		tail = tail.Next
	}
	k = k % count
	k = (count - k) % count

	if k == 0 {
		return head
	}

	mid := head
	for i := 1; i < k; i++ {
		mid = mid.Next
	}

	res := mid.Next
	tail.Next, mid.Next = head, nil

	return res
}
