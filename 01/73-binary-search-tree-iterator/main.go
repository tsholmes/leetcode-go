package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	vals []int
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{}
	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		b.vals = append(b.vals, n.Val)
		walk(n.Right)
	}
	walk(root)
	return b
}

/** @return the next smallest number */
func (b *BSTIterator) Next() int {
	var v int
	v, b.vals = b.vals[0], b.vals[1:]
	return v
}

/** @return whether we have a next smallest number */
func (b *BSTIterator) HasNext() bool {
	return len(b.vals) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
