package main

func main() {

}

type TwoSum struct {
	nums map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		nums: map[int]int{},
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for k := range this.nums {
		k2 := value - k
		if k == k2 && this.nums[k] > 1 {
			return true
		} else if k != k2 && this.nums[k2] > 0 {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
