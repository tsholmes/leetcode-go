package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var d1 []int
	for l1 != nil {
		d1 = append(d1, l1.Val)
		l1 = l1.Next
	}
	for i := 0; i < len(d1)/2; i++ {
		j := len(d1) - i - 1
		d1[i], d1[j] = d1[j], d1[i]
	}
	var d2 []int
	for l2 != nil {
		d2 = append(d2, l2.Val)
		l2 = l2.Next
	}
	for i := 0; i < len(d2)/2; i++ {
		j := len(d2) - i - 1
		d2[i], d2[j] = d2[j], d2[i]
	}

	carry := 0
	var res []int
	for i := 0; i < len(d1) || i < len(d2); i++ {
		var n1, n2 int
		if i < len(d1) {
			n1 = d1[i]
		}
		if i < len(d2) {
			n2 = d2[i]
		}
		d := n1 + n2 + carry
		res = append(res, d%10)
		carry = d / 10
	}
	if carry > 0 {
		res = append(res, carry)
	}

	var rn *ListNode
	for _, n := range res {
		rn = &ListNode{Val: n, Next: rn}
	}

	return rn
}
