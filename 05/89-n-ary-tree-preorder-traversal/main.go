package main

func main() {

}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		for _, c := range n.Children {
			walk(c)
		}
	}
	walk(root)
	return res
}
