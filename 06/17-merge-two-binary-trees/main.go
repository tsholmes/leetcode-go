package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var merge func(n1, n2 *TreeNode) *TreeNode
	merge = func(n1, n2 *TreeNode) *TreeNode {
		if n1 == nil && n2 == nil {
			return nil
		}
		val := 0
		if n1 != nil {
			val += n1.Val
		} else {
			n1 = &TreeNode{}
		}
		if n2 != nil {
			val += n2.Val
		} else {
			n2 = &TreeNode{}
		}

		return &TreeNode{
			Val:   val,
			Left:  merge(n1.Left, n2.Left),
			Right: merge(n1.Right, n2.Right),
		}
	}
	return merge(t1, t2)
}
