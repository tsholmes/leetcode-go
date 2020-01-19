package main

import "math/rand"

func main() {

}

type Solution struct {
	nums []int
	orig []int
}

func Constructor(nums []int) Solution {
	orig := make([]int, len(nums))
	copy(orig, nums)
	return Solution{
		nums: nums,
		orig: orig,
	}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	copy(s.nums, s.orig)
	return s.nums
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	for j := len(s.nums) - 1; j >= 1; j-- {
		i := rand.Intn(j + 1)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
	return s.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
