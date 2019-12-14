package main

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxIn := map[int]int{}
	for i, v := range inorder {
		idxIn[v] = i
	}
	var build func(int, int, int) *TreeNode
	build = func(postI, inL, inR int) *TreeNode {
		if inL == inR || postI == -1 {
			return nil
		}
		inI := idxIn[postorder[postI]]
		postI2 := postI - (inR - inI)
		return &TreeNode{
			Val:   postorder[postI],
			Right: build(postI-1, inI+1, inR),
			Left:  build(postI2, inL, inI),
		}
	}
	return build(len(postorder)-1, 0, len(inorder))
}
