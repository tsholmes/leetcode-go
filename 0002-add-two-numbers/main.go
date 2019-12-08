package main

import "fmt"

func main() {
	fmt.Println(flatten(addTwoNumbers(
		digitsToList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		digitsToList([]int{5, 6, 4}),
	)))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func digitsToList(digits []int) *ListNode {
	if len(digits) == 0 {
		return nil
	}
	return &ListNode{
		Val:  digits[0],
		Next: digitsToList(digits[1:]),
	}
}

func flatten(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1 := flatten(l1)
	d2 := flatten(l2)

	var d []int

	carry := 0
	for i := 0; i < len(d1) || i < len(d2); i++ {
		var v1, v2 int
		if i < len(d1) {
			v1 = d1[i]
		}
		if i < len(d2) {
			v2 = d2[i]
		}
		v := v1 + v2 + carry
		carry = v / 10
		v = v % 10
		d = append(d, v)
	}
	if carry != 0 {
		d = append(d, carry)
	}

	return digitsToList(d)
}
