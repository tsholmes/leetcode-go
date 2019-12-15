package main

import "fmt"

func main() {
	fmt.Println(getDecimalValue())
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	val := 0
	for head != nil {
		val = val * 2
		val = val | head.Val
		head = head.Next
	}
	return val
}
