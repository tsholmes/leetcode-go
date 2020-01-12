package main

func main() {

}

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger interface {
	IsInteger() bool
	GetInteger() int
	SetInteger(value int)
	Add(elem NestedInteger)
	GetList() []NestedInteger
}

func depthSum(nestedList []NestedInteger) int {
	var sum int
	var search func(n NestedInteger, depth int)
	search = func(n NestedInteger, depth int) {
		if n.IsInteger() {
			sum += n.GetInteger() * depth
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
