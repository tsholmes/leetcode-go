package main

import (
	"fmt"
	"time"
)

func main() {
	ns := make([]int, 10000)
	for i := range ns {
		ns[i] = 10000 - i
	}
	start := time.Now()
	res := nextGreaterElements(ns)
	end := time.Now()
	fmt.Println(end.Sub(start))
	for i, n := range res {
		exp := 10000
		if i == 0 {
			exp = -1
		}
		if n != exp {
			panic(fmt.Sprint(i, n, exp))
		}
	}
}

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i, n := range nums {
		next := -1
		j := (i + 1) % len(nums)
		for j != i {
			if nums[j] > n {
				next = nums[j]
				break
			}
			j++
			if j == len(nums) {
				j = 0
			}
		}
		res[i] = next
	}
	return res
}
