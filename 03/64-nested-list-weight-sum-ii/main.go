package main

func main() {

}

type NestedInteger interface {
	IsInteger() bool
	GetInteger() int
	SetInteger(value int)
	Add(elem NestedInteger)
	GetList() []NestedInteger
}

func depthSumInverse(nestedList []NestedInteger) int {
	var maxDepth int
	var search func(n NestedInteger, depth int)
	search = func(n NestedInteger, depth int) {
		if depth > maxDepth {
			maxDepth = depth
		}
		if n.IsInteger() {
			return
		}
		for _, n2 := range n.GetList() {
			search(n2, depth+1)
		}
	}
	for _, n := range nestedList {
		search(n, 1)
	}

	var sum int
	search = func(n NestedInteger, depth int) {
		if n.IsInteger() {
			sum += n.GetInteger() * (maxDepth + 1 - depth)
			return
		}
		for _, n2 := range n.GetList() {
			search(n2, depth+1)
		}
	}
	for _, n := range nestedList {
		search(n, 1)
	}

	return sum
}
