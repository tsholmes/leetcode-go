package main

func main() {

}

type MyStack struct {
	vals []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	nvals := make([]int, 0, len(s.vals)+1)
	nvals = append(nvals, x)
	nvals = append(nvals, s.vals...)
	s.vals = nvals
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	v := s.vals[0]
	s.vals = s.vals[1:]
	return v
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.vals[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.vals) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
