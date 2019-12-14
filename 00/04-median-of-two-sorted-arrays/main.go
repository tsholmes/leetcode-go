package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 4}, []int{3, 5}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 6, 7, 9}, []int{2, 4, 5, 8, 10}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 6, 7, 9}, []int{2, 4, 5, 8}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 5}, []int{}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(append([]int{}, nums1...), nums2...)
	sort.Ints(nums)

	mid := len(nums) / 2
	if len(nums)%2 == 1 {
		return float64(nums[mid])
	} else {
		return (float64(nums[mid]) + float64(nums[mid-1])) / 2.0
	}
}
