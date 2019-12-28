package main

import "fmt"

func main() {
	do := func(nums []int, k int) {
		rotate(nums, k)
		fmt.Println(nums)
	}

	do([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	do([]int{-1, 100, 3, 99}, 2)
}

func rotate(nums []int, k int) {
	n2 := make([]int, len(nums))
	copy(n2, nums)
	for i := range nums {
		nums[(i+k)%len(nums)] = n2[i]
	}
}
