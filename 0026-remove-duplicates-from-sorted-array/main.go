package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := nums[0]
	count := 1
	for _, v := range nums[1:] {
		if v != last {
			nums[count] = v
			last = v
			count++
		}
	}
	return count
}
