package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(jump([]int{1, 1, 7, 9, 1, 1, 1, 1, 1, 1, 1}))
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	ends := []int{nums[0]}
	jumpPos := 0
	for i := 1; i < len(nums); i++ {
		for ends[jumpPos] < i {
			jumpPos++
		}
		if len(ends) <= jumpPos+1 {
			ends = append(ends, i+nums[i])
		} else if ends[jumpPos+1] < i+nums[i] {
			ends[jumpPos+1] = i + nums[i]
		}
	}
	return jumpPos + 1
}
