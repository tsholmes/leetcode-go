package main

func main() {

}

type MovingAverage struct {
	nums []int
	i    int
	sum  int
	full bool
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{nums: make([]int, size)}
}

func (m *MovingAverage) Next(val int) float64 {
	m.sum += val - m.nums[m.i]
	m.nums[m.i] = val
	m.i = (m.i + 1) % len(m.nums)
	if m.i == 0 {
		m.full = true
	}
	if m.full {
		return float64(m.sum) / float64(len(m.nums))
	} else {
		return float64(m.sum) / float64(m.i)
	}
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
