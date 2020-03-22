package main

func main() {

}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		for _, c := range n.Children {
			walk(c)
		}
		res = append(res, n.Val)
	}
	walk(root)
	return res
}
