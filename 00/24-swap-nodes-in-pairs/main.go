package main

import "fmt"

func main() {
	fmt.Println(flatten(swapPairs(unflatten(1, 2, 3, 4))))
	fmt.Println(flatten(swapPairs(unflatten(1, 2, 3, 4, 5))))
	fmt.Println(flatten(swapPairs(unflatten(1))))
	fmt.Println(flatten(swapPairs(unflatten())))
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var b *ListNode
	n := head

	for {
		if b != nil {
			b.Next = n.Next
		} else {
			head = n.Next
		}
		n.Next.Next, n.Next = n, n.Next.Next
		b = n
		n = b.Next
		if n == nil || n.Next == nil {
			break
		}
	}

	return head
}
