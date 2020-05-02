package main

import "sort"

func main() {

}

func findDiagonalOrder(nums [][]int) []int {
	var res []int

	var order [][2]int
	for i, row := range nums {
		for j := range row {
			order = append(order, [2]int{i, j})
		}
	}

	sort.Slice(order, func(i, j int) bool {
		a, b := order[i], order[j]
		if a[0]+a[1] < b[0]+b[1] {
			return true
		} else if a[0]+a[1] > b[0]+b[1] {
			return false
		}
		return a[0] > b[0]
	})

	for _, ij := range order {
		res = append(res, nums[ij[0]][ij[1]])
	}

	return res
}
