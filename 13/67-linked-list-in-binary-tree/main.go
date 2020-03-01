package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	type state struct {
		ln *ListNode
		tn *TreeNode
	}
	var queue []state
	seen := map[state]bool{}

	var find func(n *TreeNode)
	find = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Val == head.Val {
			queue = append(queue, state{
				ln: head,
				tn: n,
			})
		}
		find(n.Left)
		find(n.Right)
	}
	find(root)

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if s.ln == nil {
			return true
		}
		if s.tn == nil {
			continue
		}
		if s.ln.Val != s.tn.Val {
			continue
		}
		if seen[s] {
			continue
		}
		seen[s] = true
		queue = append(queue, state{
			ln: s.ln.Next,
			tn: s.tn.Left,
		}, state{
			ln: s.ln.Next,
			tn: s.tn.Right,
		})
	}

	return false
}
