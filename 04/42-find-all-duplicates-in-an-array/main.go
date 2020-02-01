package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n != i+1 && nums[n-1] != n {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			i--
		}
	}
	count := 0
	for i, n := range nums {
		if n != i+1 {
			nums[count] = n
			count++
		}
	}
	nums = nums[:count]
	return nums
}
