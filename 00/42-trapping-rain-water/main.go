package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 3}))
}

func trap(height []int) int {
	as := make([]int, len(height))
	for i := 0; i < len(height)-1; {
		j := i + 1
		for j < len(height) && height[j] < height[i] {
			j++
		}
		if j == len(height) {
			i++
			continue
		}
		for k := i + 1; k < j; k++ {
			a := height[i] - height[k]
			if a > as[k] {
				as[k] = a
			}
		}
		i = j
	}
	for i := len(height) - 1; i > 0; {
		j := i - 1
		for j >= 0 && height[j] < height[i] {
			j--
		}
		if j == -1 {
			i--
			continue
		}
		for k := i - 1; k > j; k-- {
			a := height[i] - height[k]
			if a > as[k] {
				as[k] = a
			}
		}
		i = j
	}
	area := 0
	for _, a := range as {
		area += a
	}
	return area
}
