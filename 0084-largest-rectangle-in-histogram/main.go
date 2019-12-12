package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{1}))
	fmt.Println(largestRectangleArea([]int{0, 3, 0, 1, 0}))
}

func largestRectangleArea(heights []int) int {
	max := 0
	for i := range heights {
		min := heights[i]
		for j := i; j >= 0; j-- {
			if heights[j] < min {
				min = heights[j]
			}
			a := min * (i - j + 1)
			if a > max {
				max = a
			}
		}
	}
	return max
}
