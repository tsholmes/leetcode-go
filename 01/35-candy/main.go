package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}

func candy(ratings []int) int {
	height := make([]int, len(ratings))
	for i := range height {
		if height[i] != 0 {
			continue
		}
		height[i] = 1
		for j := i - 1; j >= 0; j-- {
			if ratings[j] <= ratings[j+1] {
				break
			}
			if height[j] < height[i]+(i-j) {
				height[j] = height[i] + (i - j)
			}
		}
		for j := i + 1; j < len(height); j++ {
			if ratings[j] <= ratings[j-1] {
				break
			}
			if height[j] < height[i]+(j-i) {
				height[j] = height[i] + (j - i)
			}
		}
	}
	sum := 0
	for _, v := range height {
		sum += v
	}
	return sum
}
