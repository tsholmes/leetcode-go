package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	s := Constructor([]int{3, 3, 3, 2, 1})
	fmt.Println(s.Pick(1))
	fmt.Println(s.Pick(2))
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(3))
}

type Solution struct {
	nums []int
	idxs []int
}

func Constructor(nums []int) Solution {
	idxs := make([]int, len(nums))
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return nums[idxs[i]] < nums[idxs[j]]
	})
	return Solution{nums, idxs}
}

func (s *Solution) Pick(target int) int {
	r := sort.Search(len(s.nums), func(i int) bool { return s.nums[s.idxs[i]] > target }) - 1
	l := len(s.nums) - sort.Search(len(s.nums), func(i int) bool { return s.nums[s.idxs[len(s.nums)-i-1]] < target })
	return s.idxs[rand.Intn(r+1-l)+l]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
