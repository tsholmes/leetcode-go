package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	idxPre := map[int]int{}
	for i, v := range preorder {
		idxPre[v] = i
	}
	idxIn := map[int]int{}
	for i, v := range inorder {
		idxIn[v] = i
	}
	var build func(int, int, int) *TreeNode
	build = func(preI, inL, inR int) *TreeNode {
		if inL == inR || preI == len(preorder) {
			return nil
		}
		inI := idxIn[preorder[preI]]
		preI2 := preI + inI - inL + 1
		return &TreeNode{
			Val:   preorder[preI],
			Left:  build(preI+1, inL, inI),
			Right: build(preI2, inI+1, inR),
		}
	}
	return build(0, 0, len(inorder))
}
