package main

func main() {

}

type MinStack [][2]int

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(*this) > 0 {
		v := (*this)[len(*this)-1][1]
		if v < min {
			min = v
		}
	}
	*this = append(*this, [2]int{x, min})
}

func (this *MinStack) Pop() {
	if l := len(*this); l > 0 {
		*this = (*this)[:l-1]
	}
}

func (this *MinStack) Top() int {
	if l := len(*this); l > 0 {
		return (*this)[l-1][0]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if l := len(*this); l > 0 {
		return (*this)[l-1][1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
