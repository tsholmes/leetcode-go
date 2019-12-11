package main

import "fmt"

func main() {
	do := func(nums ...int) {
		sortColors(nums)
		fmt.Println(nums)
	}

	do(2, 0, 2, 1, 1, 0)
	do(1, 2, 0)
}

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	start := 0
	end := len(nums) - 1
	for i := 0; i <= end; i++ {
		switch nums[i] {
		case 0:
			nums[i], nums[start] = nums[start], nums[i]
			start++
		case 2:
			nums[i], nums[end] = nums[end], nums[i]
			end--
			i--
		}
	}
}
