package main

import "fmt"

func main() {
	fmt.Println(flatten(mergeKLists([]*ListNode{
		unflatten(1, 4, 5),
		unflatten(1, 3, 4),
		unflatten(2, 6),
	})))
	fmt.Println(flatten(mergeKLists([]*ListNode{
		unflatten(),
	})))
	fmt.Println(flatten(mergeKLists([]*ListNode{
		unflatten(),
		unflatten(1),
		unflatten(),
	})))
	fmt.Println(flatten(mergeKLists([]*ListNode{})))
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

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	headi := -1
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		if head == nil || lists[i].Val < head.Val {
			head = lists[i]
			headi = i
		}
	}
	if head == nil {
		return nil
	}
	lists[headi] = lists[headi].Next

	tail := head
	for {
		count := 0
		for _, l := range lists {
			if l != nil {
				lists[count] = l
				count++
			}
		}
		lists = lists[:count]
		if count == 0 {
			break
		}

		mini := 0
		for i := 1; i < len(lists); i++ {
			if lists[i].Val < lists[mini].Val {
				mini = i
			}
		}

		tail.Next = lists[mini]
		tail = tail.Next
		lists[mini] = lists[mini].Next
	}

	return head
}
