package main

import "fmt"

func main() {
	do := func(nums ...int) {
		l := unflatten(nums...)
		reorderList(l)
		fmt.Println(flatten(l))
	}
	do(1, 2, 3, 4)
	do(1, 2, 3, 4, 5)
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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	for i, n := range nodes {
		j := len(nodes) - i
		if i < (len(nodes)+1)/2 {
			j--
		}
		if i == j {
			n.Next = nil
		} else {
			n.Next = nodes[j]
		}
	}
}
