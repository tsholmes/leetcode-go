package main

import "strconv"

func main() {

}

func findContestMatch(n int) string {
	type node struct {
		val  int
		l, r *node
	}
	var nodes []node
	for i := 1; i <= n; i++ {
		nodes = append(nodes, node{val: i})
	}
	for len(nodes) > 1 {
		var nnodes []node
		for i := 0; i < len(nodes)/2; i++ {
			ni, nj := nodes[i], nodes[len(nodes)-i-1]
			nnodes = append(nnodes, node{l: &ni, r: &nj})
		}
		nodes = nnodes
	}
	var work []byte
	var build func(*node)
	build = func(n *node) {
		if n.val > 0 {
			work = strconv.AppendInt(work, int64(n.val), 10)
			return
		}
		work = append(work, '(')
		build(n.l)
		work = append(work, ',')
		build(n.r)
		work = append(work, ')')
	}
	build(&nodes[0])
	return string(work)
}
