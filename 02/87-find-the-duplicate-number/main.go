package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 3, 3, 3}))
	fmt.Println(findDuplicate([]int{2, 2, 2}))
	fmt.Println(findDuplicate([]int{1, 1}))
}

func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, n := range nums {
			if n >= left && n <= mid {
				count++
			}
		}
		if count <= (mid-left)+1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
