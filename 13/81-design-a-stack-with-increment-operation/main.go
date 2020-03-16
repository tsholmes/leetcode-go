package main

func main() {

}

type CustomStack struct {
	els []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		els: make([]int, 0, maxSize),
	}
}

func (c *CustomStack) Push(x int) {
	if len(c.els) < cap(c.els) {
		c.els = append(c.els, x)
	}
}

func (c *CustomStack) Pop() int {
	if len(c.els) == 0 {
		return -1
	}
	v := c.els[len(c.els)-1]
	c.els = c.els[:len(c.els)-1]
	return v
}

func (c *CustomStack) Increment(k int, val int) {
	if k > len(c.els) {
		k = len(c.els)
	}
	for i := range c.els[:k] {
		c.els[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
