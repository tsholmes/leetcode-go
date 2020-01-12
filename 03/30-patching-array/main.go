package main

import "fmt"

func main() {
	fmt.Println(minPatches([]int{1, 3}, 6))
	fmt.Println(minPatches([]int{1, 5, 10}, 20))
	fmt.Println(minPatches([]int{1, 2, 2}, 5))
	fmt.Println(minPatches([]int{}, 7))
}

func minPatches(nums []int, n int) int {
	max := 0
	count := 0
	for _, num := range nums {
		for max < num-1 && max < n {
			count++
			max = max*2 + 1
		}
		if max >= n {
			break
		}
		max += num
	}
	for max < n {
		count++
		max = max*2 + 1
	}
	return count
}
