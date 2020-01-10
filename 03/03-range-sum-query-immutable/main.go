package main

func main() {

}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	for i, n := range nums {
		sums[i] = n
		if i > 0 {
			sums[i] += sums[i-1]
		}
	}
	return NumArray{sums: sums}
}

func (n *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return n.sums[j]
	}
	return n.sums[j] - n.sums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
