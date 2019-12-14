package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	work := make([]int, len(nums))
	var search func(int, int)
	search = func(i int, j int) {
		if j == len(nums) {
			return
		}
		work[i] = nums[j]
		res = append(res, append([]int{}, work[:i+1]...))
		search(i+1, j+1)
		search(i, j+1)
	}
	search(0, 0)
	return res
}
