package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			min := height[i]
			if height[j] < min {
				min = height[j]
			}
			a := min * (j - i)
			if a > max {
				max = a
			}
		}
	}
	return max
}
