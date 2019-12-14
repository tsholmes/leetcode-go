package main

import "fmt"

func main() {
	fmt.Println(flatten(reverseKGroup(unflatten(1, 2, 3, 4, 5), 2)))
	fmt.Println(flatten(reverseKGroup(unflatten(1, 2, 3, 4, 5), 3)))
	fmt.Println(flatten(reverseKGroup(unflatten(), 2)))
	fmt.Println(flatten(reverseKGroup(unflatten(1, 2), 2)))
	fmt.Println(flatten(reverseKGroup(unflatten(1, 2, 3, 4), 4)))
	fmt.Println(flatten(reverseKGroup(unflatten(1, 2, 3, 4, 5, 6), 3)))
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}
	swap := func(b1 *ListNode, b2 *ListNode) {
		if b1 == nil {
			if b2 == head {
				head, b2.Next, b2.Next.Next = b2.Next, b2.Next.Next, b2
			} else {
				head, head.Next, b2.Next, b2.Next.Next = b2.Next, b2.Next.Next, head, head.Next
			}
		} else {
			if b1.Next == b2 {
				b1.Next, b2.Next, b2.Next.Next = b2.Next, b2.Next.Next, b2
			} else {
				b1.Next, b1.Next.Next, b2.Next, b2.Next.Next = b2.Next, b2.Next.Next, b1.Next, b1.Next.Next
			}
		}
	}

	var b *ListNode
	bs := make([]*ListNode, 0, k)

	for {
		bs = bs[:0]
		bs = append(bs, b)
		for i := 1; i < k; i++ {
			if b == nil {
				b = head
			} else {
				b = b.Next
			}
			if b.Next == nil {
				break
			}
			bs = append(bs, b)
		}

		if len(bs) == k {
			b = bs[1]
			for i := 0; i < k/2; i++ {
				j := k - i - 1
				b1, b2 := bs[i], bs[j]
				bs[i+1] = bs[j].Next
				swap(b1, b2)
			}
		}
		if b.Next == nil {
			break
		}
	}

	return head
}
