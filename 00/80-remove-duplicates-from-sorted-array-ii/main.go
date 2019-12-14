package main

import "fmt"

func main() {
	do := func(nums ...int) {
		count := removeDuplicates(nums)
		fmt.Println(nums[:count])
	}
	do(1, 1, 1, 2, 2, 3)
	do(0, 0, 1, 1, 1, 1, 2, 3, 3)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[count-1] || nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
