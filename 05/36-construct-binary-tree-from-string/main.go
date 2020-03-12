package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {
	if s == "" {
		return nil
	}
	var build func() *TreeNode
	build = func() *TreeNode {
		val := 0
		neg := false
		if s[0] == '-' {
			neg = true
			s = s[1:]
		}
		for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
			val *= 10
			val += int(s[0] - '0')
			s = s[1:]
		}
		if neg {
			val = -val
		}
		var left, right *TreeNode
		if len(s) > 0 && s[0] == '(' {
			s = s[1:]
			left = build()
			s = s[1:]
		}
		if len(s) > 0 && s[0] == '(' {
			s = s[1:]
			right = build()
			s = s[1:]
		}
		return &TreeNode{
			Val:   val,
			Left:  left,
			Right: right,
		}
	}
	return build()
}
