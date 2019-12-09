package main

import "fmt"

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func threeSumClosest(nums []int, target int) int {
	seen := map[int]bool{}

	min := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			s := nums[i] + nums[j]
			if seen[s] {
				continue
			}
			seen[s] = true
			for k := j + 1; k < len(nums); k++ {
				v := s + nums[k]
				if abs(target-v) < abs(target-min) {
					min = v
				}
			}
		}
	}

	return min
}
