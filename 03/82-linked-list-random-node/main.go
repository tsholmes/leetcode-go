package main

import "math/rand"

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head}
}

/** Returns a random node's value. */
func (s *Solution) GetRandom() int {
	n := s.head
	val := 0
	count := 0
	for n != nil {
		count++
		if rand.Intn(count) == 0 {
			val = n.Val
		}
		n = n.Next
	}
	return val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
