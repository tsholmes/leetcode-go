package main

import "sort"

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for _, row := range matrix {
		idx := sort.Search(len(row), func(i int) bool {
			return row[i] >= target
		})
		if idx < len(row) && row[idx] == target {
			return true
		}
	}
	return false
}
