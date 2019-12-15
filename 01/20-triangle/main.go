package main

import "fmt"

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	mins := make([]int, len(triangle))
	mins[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			if j == i {
				mins[j] = triangle[i][j] + mins[j-1]
			} else if j == 0 {
				mins[j] += triangle[i][j]
			} else if mins[j-1] < mins[j] {
				mins[j] = mins[j-1] + triangle[i][j]
			} else {
				mins[j] += triangle[i][j]
			}
		}
	}
	min := mins[0]
	for _, v := range mins {
		if v < min {
			min = v
		}
	}
	return min
}
