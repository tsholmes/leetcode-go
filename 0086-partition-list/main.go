package main

import "fmt"

func main() {
	fmt.Println(flatten(partition(unflatten(1, 4, 3, 2, 5, 2), 3)))
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

func partition(head *ListNode, x int) *ListNode {
	var preHead, preTail, postHead, postTail *ListNode
	for head != nil {
		n := head
		head = head.Next
		n.Next = nil
		if n.Val < x {
			if preHead == nil {
				preHead, preTail = n, n
			} else {
				preTail.Next = n
				preTail = n
			}
		} else {
			if postHead == nil {
				postHead, postTail = n, n
			} else {
				postTail.Next = n
				postTail = n
			}
		}
	}
	if preHead == nil {
		return postHead
	}
	preTail.Next = postHead
	return preHead
}
